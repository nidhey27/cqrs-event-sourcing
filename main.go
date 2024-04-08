package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
	"github.com/nidhey27/cqrs-go/ent"
	"github.com/nidhey27/cqrs-go/ent/account"
	"github.com/nidhey27/cqrs-go/ent/transaction"
	"github.com/nidhey27/cqrs-go/protobuf/account_command"
	"github.com/nidhey27/cqrs-go/protobuf/account_query"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type AccountCommandService struct {
	dbClient *ent.Client
	account_command.UnimplementedAccountCommandServiceServer
}

type AccountQueryService struct {
	dbClient *ent.Client
	account_query.UnimplementedAccountQueryServiceServer
}

func main() {

	client, err := ent.Open("postgres", "host=localhost port=5433 user=nidhey dbname=bank password=password sslmode=disable")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	listener, err := net.Listen("tcp", ":9000")
	if err != nil {
		panic(err)
	}

	server := grpc.NewServer()

	account_command.RegisterAccountCommandServiceServer(server, &AccountCommandService{
		dbClient: client,
	})
	account_query.RegisterAccountQueryServiceServer(server, &AccountQueryService{
		dbClient: client,
	})

	reflection.Register(server)
	log.Println("Server started at PORT 💻: 9000")
	if err := server.Serve(listener); err != nil {
		panic(err)
	}
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
	go func(acs *AccountCommandService, account_no string) {
		ctx := context.Background()
		log.Printf("Calculating balance of %v after withdrawal...", account_no)
		balance, err := CalucateBalance(ctx, acs.dbClient, account_no)
		if err != nil {
			log.Println(err)
			return
		}
		_, err = acs.dbClient.Account.Update().
			SetAccountBalance(balance).
			Where(account.AccountNumber(account_no)).Save(ctx)
		if err != nil {
			log.Println(err)
			return
		}

		log.Printf("Account Balance Updated: %v:%v", account_no, balance)
	}(acs, req.AccountNo)
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
	go func(acs *AccountCommandService, account_no string) {
		ctx := context.Background()
		log.Printf("Calculating balance of %v after deposit...", account_no)
		balance, err := CalucateBalance(ctx, acs.dbClient, account_no)
		if err != nil {
			log.Println(err)
			return
		}
		_, err = acs.dbClient.Account.Update().
			SetAccountBalance(balance).
			Where(account.AccountNumber(account_no)).Save(ctx)
		if err != nil {
			log.Println(err)
			return
		}

		log.Printf("Account Balance Updated: %v:%v", account_no, balance)
	}(acs, req.AccountNo)
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
	go func(acs *AccountCommandService, account_no string) {
		ctx := context.Background()
		log.Printf("Calculating balance of %v after withdrawal...", account_no)
		balance, err := CalucateBalance(ctx, acs.dbClient, account_no)
		if err != nil {
			log.Println(err)
			return
		}
		_, err = acs.dbClient.Account.Update().
			SetAccountBalance(balance).
			Where(account.AccountNumber(account_no)).Save(ctx)
		if err != nil {
			log.Println(err)
			return
		}

		log.Printf("Account Balance Updated: %v:%v", account_no, balance)
	}(acs, req.AccountFrom)
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
	go func(acs *AccountCommandService, account_no string) {
		ctx := context.Background()
		log.Printf("Calculating balance of %v after deposit...", account_no)
		balance, err := CalucateBalance(ctx, acs.dbClient, account_no)
		if err != nil {
			log.Println(err)
			return
		}
		_, err = acs.dbClient.Account.Update().
			SetAccountBalance(balance).
			Where(account.AccountNumber(account_no)).Save(ctx)
		if err != nil {
			log.Println(err)
			return
		}

		log.Printf("Account Balance Updated: %v:%v", account_no, balance)
	}(acs, req.AccountTp)
	// Calculate Account Balance END

	return &account_command.TransferResponse{
		Message: fmt.Sprintf("INR %v transfered from %v to %v", req.Amount, req.AccountFrom, req.AccountTp),
	}, nil
}

func CalucateBalance(ctx context.Context, dbClient *ent.Client, accountNo string) (float64, error) {
	transactions := dbClient.Transaction.Query().
		Where(transaction.AccountNumber(accountNo)).AllX(ctx)
	var balance float64 = 0
	for _, tx := range transactions {
		switch tx.Type {
		case "credit":
			balance += tx.Amount
		case "debit":
			balance -= tx.Amount
		}
	}

	return balance, nil
}

func (aqs *AccountQueryService) GetBalance(ctx context.Context, req *account_query.BalanceRequest) (*account_query.BalanceResponse, error) {

	transactions := aqs.dbClient.Account.Query().
		Where(account.AccountNumber(req.AccountNo)).OnlyX(ctx)

	return &account_query.BalanceResponse{
		Balance: float32(transactions.AccountBalance),
	}, nil

}
