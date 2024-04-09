package activities

import (
	"context"

	"github.com/nidhey27/cqrs-go/ent"
)

func GetBalance(ctx context.Context, transactions []*ent.Transaction) (float64, error) {
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
