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
	// Start a database transaction
	tx, err := acs.dbClient.BeginTx(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to begin transaction: %v", err)
	}
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			log.Printf("recovered panic: %v", p)
		} else if err != nil {
			tx.Rollback()
			log.Printf("rolling back transaction due to error: %v", err)
		}
	}()

	txn_id := uuid.New()

	// Debit from account
	if _, err := acs.debit(ctx, txn_id.String(), req.AccountNo, float64(req.Amount)); err != nil {
		log.Fatalf("failed to withdraw amount: %v", err)
		return &account_command.WithdrawResponse{
			Message: "Something went wrong! Please try again",
		}, err
	}

	// Calculate Account Balance
	if err := acs.calculateBalance(ctx, "account_workflow_debit", req.AccountNo); err != nil {
		return nil, err
	}

	// Commit the transaction
	if err := tx.Commit(); err != nil {
		return nil, fmt.Errorf("failed to commit transaction: %v", err)
	}

	return &account_command.WithdrawResponse{
		Message: "Withdrawal Successful!",
		Balance: req.Amount,
	}, nil
}

func (acs *AccountCommandService) Deposite(ctx context.Context, req *account_command.DespositeRequst) (*account_command.DepositeResponse, error) {
	// Start a database transaction
	tx, err := acs.dbClient.BeginTx(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to begin transaction: %v", err)
	}
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			log.Printf("recovered panic: %v", p)
		} else if err != nil {
			tx.Rollback()
			log.Printf("rolling back transaction due to error: %v", err)
		}
	}()

	txn_id := uuid.New()

	// Credit to account
	if _, err := acs.credit(ctx, txn_id.String(), req.AccountNo, float64(req.Amount)); err != nil {
		log.Fatalf("failed to deposit amount: %v", err)
		return &account_command.DepositeResponse{
			Message: "Something went wrong! Please try again",
		}, err
	}

	// Calculate Account Balance
	if err := acs.calculateBalance(ctx, "account_workflow_credit", req.AccountNo); err != nil {
		return nil, err
	}

	// Commit the transaction
	if err := tx.Commit(); err != nil {
		return nil, fmt.Errorf("failed to commit transaction: %v", err)
	}

	return &account_command.DepositeResponse{
		Message: "Deposit Successful!",
		Balance: req.Amount,
	}, nil
}

func (acs *AccountCommandService) Transfer(ctx context.Context, req *account_command.TransferRequest) (*account_command.TransferResponse, error) {
	// Start a database transaction
	tx, err := acs.dbClient.BeginTx(ctx, nil)

	if err != nil {
		return nil, fmt.Errorf("failed to begin transaction: %v", err)
	}

	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			log.Printf("recovered panic: %v", p)
		} else if err != nil {
			tx.Rollback()
			log.Printf("rolling back transaction due to error: %v", err)
		}
	}()
	txn_id := uuid.New()

	// Debit from "FROM account"
	if _, err := acs.debit(ctx, txn_id.String(), req.AccountFrom, float64(req.Amount)); err != nil {
		log.Fatalf("failed to transfer(debit) amount: %v", err)
		return &account_command.TransferResponse{
			Message: "Something went wrong! Please try again",
		}, err
	}

	// Calculate Account Balance
	if err := acs.calculateBalance(ctx, "account_workflow_debit", req.AccountFrom); err != nil {
		return nil, err
	}

	// Credit to "TO account"
	if _, err := acs.credit(ctx, txn_id.String(), req.AccountTp, float64(req.Amount)); err != nil {
		log.Fatalf("failed to transfer(credit) amount: %v", err)
		return &account_command.TransferResponse{
			Message: "Something went wrong! Please try again",
		}, err
	}

	// Calculate Account Balance
	if err := acs.calculateBalance(ctx, "account_workflow_credit", req.AccountTp); err != nil {
		return nil, err
	}

	// Commit the transaction
	if err := tx.Commit(); err != nil {
		return nil, fmt.Errorf("failed to commit transaction: %v", err)
	}

	return &account_command.TransferResponse{
		Message: fmt.Sprintf("INR %v transfered from %v to %v", req.Amount, req.AccountFrom, req.AccountTp),
	}, nil
}

func (acs *AccountCommandService) debit(ctx context.Context, txnID, accountNumber string, amount float64) (*ent.Transaction, error) {
	// Debit from account
	transaction, err := acs.dbClient.Transaction.Create().
		SetTransactionID(txnID).
		SetAccountNumber(accountNumber).
		SetAmount(amount). // negative amount indicates a debit
		SetType("debit").
		SetTimestamp(time.Now()).
		Save(ctx)

	if err != nil {
		return nil, fmt.Errorf("failed to debit amount: %v", err)
	}
	return transaction, nil
}

func (acs *AccountCommandService) credit(ctx context.Context, txnID, accountNumber string, amount float64) (*ent.Transaction, error) {
	// Credit to account
	transaction, err := acs.dbClient.Transaction.Create().
		SetTransactionID(txnID).
		SetAccountNumber(accountNumber).
		SetAmount(amount).
		SetType("credit").
		SetTimestamp(time.Now()).
		Save(ctx)

	if err != nil {
		return nil, fmt.Errorf("failed to credit amount: %v", err)
	}
	return transaction, nil
}

func (acs *AccountCommandService) calculateBalance(ctx context.Context, workflowID string, accountNumber string) error {
	// Calculate Account Balance
	_, err := acs.temporalClient.ExecuteWorkflow(ctx, temporal_client.StartWorkflowOptions{
		ID:        workflowID,
		TaskQueue: "account",
	}, workflows.CalculateBalance, accountNumber)
	if err != nil {
		return fmt.Errorf("unable to process request for %v: %v", accountNumber, err)
	}
	return nil
}

func (aqs *AccountQueryService) GetBalance(ctx context.Context, req *account_query.BalanceRequest) (*account_query.BalanceResponse, error) {

	transactions := aqs.dbClient.Account.Query().
		Where(account.AccountNumber(req.AccountNo)).OnlyX(ctx)

	return &account_query.BalanceResponse{
		Balance: float32(transactions.AccountBalance),
	}, nil

}
