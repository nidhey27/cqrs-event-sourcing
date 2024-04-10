package activities

import (
	"context"
	"log"

	_ "github.com/lib/pq"
	"github.com/nidhey27/cqrs-go/ent"
	"github.com/nidhey27/cqrs-go/ent/transaction"
)

func GetTransactions(ctx context.Context, account_no string) ([]*ent.Transaction, error) {

	client, err := ent.Open("postgres", "host=localhost port=5433 user=nidhey dbname=bank password=password sslmode=disable")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()

	transactions := client.Transaction.Query().
		Where(transaction.AccountNumber(account_no)).AllX(ctx)
	return transactions, nil
}
