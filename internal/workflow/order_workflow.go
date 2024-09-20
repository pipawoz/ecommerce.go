package workflow

import (
	"time"

	"go.temporal.io/sdk/workflow"
	"github.com/pipawoz/ecommerce.go/internal/db"
)

func OrderWorkflow(ctx workflow.Context, order db.Order) error {
	logger := workflow.GetLogger(ctx)
	logger.Info("OrderWorkflow started", "OrderID", order.ID)

	// Simulate order processing
	err := workflow.Sleep(ctx, 5*time.Second)
	if err != nil {
		return err
	}

	// Update order status
	err = workflow.ExecuteActivity(ctx, UpdateOrderStatus, workflow.ActivityOptions{
		StartToCloseTimeout: time.Minute,
	}, order.ID, "completed").Get(ctx, nil)
	if err != nil {
		return err
	}

	logger.Info("OrderWorkflow completed", "OrderID", order.ID)
	return nil
}

func UpdateOrderStatus(ctx workflow.Context, orderID int64, status string) error {
	// TODO: Implement database update
	return nil
}