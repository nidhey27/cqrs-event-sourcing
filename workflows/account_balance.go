package workflows

import (
	"time"

	"github.com/nidhey27/cqrs-go/activities"
	"github.com/nidhey27/cqrs-go/ent"
	"go.temporal.io/sdk/workflow"
)

func CalculateBalance(ctx workflow.Context, accountNo string) error {

	options := workflow.ActivityOptions{
		StartToCloseTimeout: time.Minute * 10,
	}

	ctx = workflow.WithActivityOptions(ctx, options)

	txnsFuture := workflow.ExecuteActivity(ctx, activities.GetTransactions, accountNo)

	var txns []*ent.Transaction
	if err := txnsFuture.Get(ctx, &txns); err != nil {
		return err
	}

	balanceFuture := workflow.ExecuteActivity(ctx, activities.GetBalance, txns)

	var balance float64
	if err := balanceFuture.Get(ctx, &balance); err != nil {
		return err
	}

	setBalanceFuture := workflow.ExecuteActivity(ctx, activities.SetBalance, balance, accountNo)
	if err := setBalanceFuture.Get(ctx, nil); err != nil {
		return err
	}

	return nil
}
