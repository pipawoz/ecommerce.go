package api

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pipawoz/ecommerce.go/internal/db"
	"github.com/pipawoz/ecommerce.go/internal/workflow"
	"go.temporal.io/sdk/client"
)

type Handler struct {
	queries    *db.Queries
	temporalClient client.Client
}

func NewHandler(queries *db.Queries, temporalClient client.Client) *Handler {
	return &Handler{
		queries:    queries,
		temporalClient: temporalClient,
	}
}

func (h *Handler) RegisterRoutes(router *gin.Engine) {
	RegisterHandlers(router, h)
}

func (h *Handler) PostOrders(c *gin.Context) {
	h.CreateOrder(c)
}

func (h *Handler) PutOrdersId(c *gin.Context, id int32) {
	h.CreateOrder(c)
}


func (h *Handler) DeleteOrdersId(c *gin.Context, id int32) {
	err := h.queries.DeleteOrder(context.Background(), id)
	if err != nil {
			if err == sql.ErrNoRows {
					c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
			} else {
					c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete order"})
			}
			return
	}

	c.Status(http.StatusNoContent)
}


func (h *Handler) GetOrders(c *gin.Context) {
	orders, err := h.queries.ListOrders(context.Background())
	if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to list orders"})
			return
	}

	c.JSON(http.StatusOK, orders)	
}


func (h *Handler) CreateOrder(c *gin.Context) {
	var req CreateOrderRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	order, err := h.queries.CreateOrder(c, db.CreateOrderParams{
		CustomerID:  req.CustomerId,
		Status:     "pending",
		TotalAmount: fmt.Sprint(req.TotalAmount),
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Start Temporal workflow
	workflowOptions := client.StartWorkflowOptions{
		ID:        "order-" + fmt.Sprint(order.ID),
		TaskQueue: "order-processing",
	}
	_, err = h.temporalClient.ExecuteWorkflow(context.Background(), workflowOptions, workflow.OrderWorkflow, order)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to start workflow"})
		return
	}

	c.JSON(http.StatusCreated, order)
}

func (h *Handler) GetOrdersId(c *gin.Context, id int32) {
	order, err := h.queries.GetOrder(context.Background(), id)
	if err != nil {
			if err == sql.ErrNoRows {
					c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
			} else {
					c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch order"})
			}
			return
	}
	
	c.JSON(http.StatusOK, order)
}


func (h *Handler) UpdateOrder(c *gin.Context, id int32) {
	var req UpdateOrderRequest
	if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
	}

	status := string(*req.Status)
	totalAmount := fmt.Sprint(req.TotalAmount)

	order, err := h.queries.UpdateOrder(context.Background(), db.UpdateOrderParams{
			ID:         id,
			Status:     status,
			TotalAmount: totalAmount,
	})
	if err != nil {
			if err == sql.ErrNoRows {
					c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
			} else {
					c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update order"})
			}
			return
	}

	c.JSON(http.StatusOK, order)
}


func (h *Handler) DeleteOrder(c *gin.Context, id int) {
	err := h.queries.DeleteOrder(context.Background(), int32(id))
	if err != nil {
			if err == sql.ErrNoRows {
					c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
			} else {
					c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete order"})
			}
			return
	}

	c.Status(http.StatusNoContent)
}
