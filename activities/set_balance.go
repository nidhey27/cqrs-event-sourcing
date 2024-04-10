package activities

import (
	"context"
	"log"

	_ "github.com/lib/pq"
	"github.com/nidhey27/cqrs-go/ent"
	"github.com/nidhey27/cqrs-go/ent/account"
)

func SetBalance(ctx context.Context, balance float64, account_no string) error {

	client, err := ent.Open("postgres", "host=localhost port=5433 user=nidhey dbname=bank password=password sslmode=disable")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()
	_, err = client.Account.Update().
		SetAccountBalance(balance).
		Where(account.AccountNumber(account_no)).Save(ctx)
	if err != nil {
		return err
	}
	return nil
}
