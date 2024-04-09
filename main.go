package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
	"github.com/nidhey27/cqrs-go/activities"
	"github.com/nidhey27/cqrs-go/ent"
	"github.com/nidhey27/cqrs-go/ent/account"
	"github.com/nidhey27/cqrs-go/protobuf/account_command"
	"github.com/nidhey27/cqrs-go/protobuf/account_query"
	"github.com/nidhey27/cqrs-go/workflows"
	temporal_client "go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type AccountCommandService struct {
	dbClient       *ent.Client
	temporalClient temporal_client.Client
	account_command.UnimplementedAccountCommandServiceServer
}

type AccountQueryService struct {
	dbClient       *ent.Client
	temporalClient temporal_client.Client
	account_query.UnimplementedAccountQueryServiceServer
}

func main() {
	// DB Setup START
	client, err := ent.Open("postgres", "host=localhost port=5433 user=nidhey dbname=bank password=password sslmode=disable")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	// DB Setup END
	//
	// Temporal Client START
	c, err := temporal_client.Dial(temporal_client.Options{})
	if err != nil {
		log.Fatalln("unable to create Temporal client", err)
	}
	defer c.Close()

	w := worker.New(c, "account_balance", worker.Options{})
	w.RegisterWorkflow(workflows.CalculateBalance)
	w.RegisterActivity(activities.GetTransactions)
	w.RegisterActivity(activities.GetBalance)
	w.RegisterActivity(activities.SetBalance)
	// Temporal Client END

	listener, err := net.Listen("tcp", ":9000")
	if err != nil {
		panic(err)
	}

	server := grpc.NewServer()

	account_command.RegisterAccountCommandServiceServer(server, &AccountCommandService{
		dbClient:       client,
		temporalClient: c,
	})
	account_query.RegisterAccountQueryServiceServer(server, &AccountQueryService{
		dbClient:       client,
		temporalClient: c,
	})

	reflection.Register(server)

	// Start the worker & gRPC Server
	log.Println("Server started at PORT 💻: 9000")
	go w.Run(worker.InterruptCh())
	go func() {
		if err := server.Serve(listener); err != nil {
			panic(err)
		}
	}()

	// Listen for termination signals
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	// Wait for termination signal
	<-stop
	w.Stop()
}

func (acs *AccountCommandService) Withdraw(ctx context.Context, req *account_command.WithdrawRequest) (*account_command.WithdrawResponse, error) {
	txn_id := uuid.New()

	_, err := acs.dbClient.Transaction.Create().
		SetTransactionID(txn_id.String()).
		SetAccountNumber(req.AccountNo).
		SetAmount(float64(req.Amount)).
		SetType("debit").
		SetTimestamp(time.Now()).
		Save(ctx)
	if err != nil {
		log.Fatalf("failed to withdraw amount: %v", err)
		return &account_command.WithdrawResponse{
			Message: "Something went wrong! Please try again",
		}, err
	}

	// Calculate Account Balance START
	_, err = acs.temporalClient.ExecuteWorkflow(context.Background(), temporal_client.StartWorkflowOptions{
		ID:        "account_workflow",
		TaskQueue: "account",
	}, workflows.CalculateBalance, req.AccountNo)
	if err != nil {
		log.Printf("Unable to process request for %v:%v", req.AccountNo, err)
	}
	// Calculate Account Balance END

	return &account_command.WithdrawResponse{
		Message: "Withdrawal Successful!",
		Balance: req.Amount,
	}, nil
}

func (acs *AccountCommandService) Deposite(ctx context.Context, req *account_command.DespositeRequst) (*account_command.DepositeResponse, error) {
	txn_id := uuid.New()

	_, err := acs.dbClient.Transaction.Create().
		SetTransactionID(txn_id.String()).
		SetAccountNumber(req.AccountNo).
		SetAmount(float64(req.Amount)).
		SetType("credit").
		SetTimestamp(time.Now()).
		Save(ctx)
	if err != nil {
		log.Fatalf("failed to deposit amount: %v", err)
		return &account_command.DepositeResponse{
			Message: "Something went wrong! Please try again",
		}, err
	}

	// Calculate Account Balance START
	_, err = acs.temporalClient.ExecuteWorkflow(context.Background(), temporal_client.StartWorkflowOptions{
		ID:        "account_workflow",
		TaskQueue: "account",
	}, workflows.CalculateBalance, req.AccountNo)
	if err != nil {
		log.Printf("Unable to process request for %v:%v", req.AccountNo, err)
	}
	// Calculate Account Balance END

	return &account_command.DepositeResponse{
		Message: "Deposit Successful!",
		Balance: req.Amount,
	}, nil
}

func (acs *AccountCommandService) Transfer(ctx context.Context, req *account_command.TransferRequest) (*account_command.TransferResponse, error) {
	// Debit from "FROM account"
	txn_id := uuid.New()

	_, err := acs.dbClient.Transaction.Create().
		SetTransactionID(txn_id.String()).
		SetAccountNumber(req.AccountFrom).
		SetAmount(float64(req.Amount)).
		SetType("debit").
		SetTimestamp(time.Now()).
		Save(ctx)

	if err != nil {
		log.Fatalf("failed to transfer(debit) amount: %v", err)
		return &account_command.TransferResponse{
			Message: "Something went wrong! Please try again",
		}, err
	}

	// Calculate Account Balance START
	_, err = acs.temporalClient.ExecuteWorkflow(context.Background(), temporal_client.StartWorkflowOptions{
		ID:        "account_workflow",
		TaskQueue: "account",
	}, workflows.CalculateBalance, req.AccountFrom)
	if err != nil {
		log.Printf("Unable to process request for %v:%v", req.AccountFrom, err)
	}
	// Calculate Account Balance END

	// Credit to "To account"
	_, err = acs.dbClient.Transaction.Create().
		SetTransactionID(txn_id.String()).
		SetAccountNumber(req.AccountTp).
		SetAmount(float64(req.Amount)).
		SetType("credit").
		SetTimestamp(time.Now()).
		Save(ctx)

	if err != nil {
		log.Fatalf("failed to transfer(credit) amount: %v", err)
		return &account_command.TransferResponse{
			Message: "Something went wrong! Please try again",
		}, err
	}

	// Calculate Account Balance START
	_, err = acs.temporalClient.ExecuteWorkflow(context.Background(), temporal_client.StartWorkflowOptions{
		ID:        "account_workflow",
		TaskQueue: "account",
	}, workflows.CalculateBalance, req.AccountTp)
	if err != nil {
		log.Printf("Unable to process request for %v:%v", req.AccountTp, err)
	}
	// Calculate Account Balance END

	return &account_command.TransferResponse{
		Message: fmt.Sprintf("INR %v transfered from %v to %v", req.Amount, req.AccountFrom, req.AccountTp),
	}, nil
}

func (aqs *AccountQueryService) GetBalance(ctx context.Context, req *account_query.BalanceRequest) (*account_query.BalanceResponse, error) {

	transactions := aqs.dbClient.Account.Query().
		Where(account.AccountNumber(req.AccountNo)).OnlyX(ctx)

	return &account_query.BalanceResponse{
		Balance: float32(transactions.AccountBalance),
	}, nil

}
