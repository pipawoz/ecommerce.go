// Package api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.16.3 DO NOT EDIT.
package api

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/oapi-codegen/runtime"
)

// Defines values for OrderStatus.
const (
	OrderStatusCancelled  OrderStatus = "cancelled"
	OrderStatusCompleted  OrderStatus = "completed"
	OrderStatusPending    OrderStatus = "pending"
	OrderStatusProcessing OrderStatus = "processing"
)

// Defines values for UpdateOrderRequestStatus.
const (
	UpdateOrderRequestStatusCancelled  UpdateOrderRequestStatus = "cancelled"
	UpdateOrderRequestStatusCompleted  UpdateOrderRequestStatus = "completed"
	UpdateOrderRequestStatusPending    UpdateOrderRequestStatus = "pending"
	UpdateOrderRequestStatusProcessing UpdateOrderRequestStatus = "processing"
)

// CreateOrderRequest defines model for CreateOrderRequest.
type CreateOrderRequest struct {
	CustomerId  int32   `json:"customer_id"`
	TotalAmount float32 `json:"total_amount"`
}

// Order defines model for Order.
type Order struct {
	CreatedAt   *time.Time   `json:"created_at,omitempty"`
	CustomerId  *int32       `json:"customer_id,omitempty"`
	Id          *int32       `json:"id,omitempty"`
	Status      *OrderStatus `json:"status,omitempty"`
	TotalAmount *float32     `json:"total_amount,omitempty"`
	UpdatedAt   *time.Time   `json:"updated_at,omitempty"`
}

// OrderStatus defines model for Order.Status.
type OrderStatus string

// UpdateOrderRequest defines model for UpdateOrderRequest.
type UpdateOrderRequest struct {
	Status      *UpdateOrderRequestStatus `json:"status,omitempty"`
	TotalAmount *float32                  `json:"total_amount,omitempty"`
}

// UpdateOrderRequestStatus defines model for UpdateOrderRequest.Status.
type UpdateOrderRequestStatus string

// PostOrdersJSONRequestBody defines body for PostOrders for application/json ContentType.
type PostOrdersJSONRequestBody = CreateOrderRequest

// PutOrdersIdJSONRequestBody defines body for PutOrdersId for application/json ContentType.
type PutOrdersIdJSONRequestBody = UpdateOrderRequest

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// List all orders
	// (GET /orders)
	GetOrders(c *gin.Context)
	// Create a new order
	// (POST /orders)
	PostOrders(c *gin.Context)
	// Delete an order
	// (DELETE /orders/{id})
	DeleteOrdersId(c *gin.Context, id int32)
	// Get an order by ID
	// (GET /orders/{id})
	GetOrdersId(c *gin.Context, id int32)
	// Update an order
	// (PUT /orders/{id})
	PutOrdersId(c *gin.Context, id int32)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
	ErrorHandler       func(*gin.Context, error, int)
}

type MiddlewareFunc func(c *gin.Context)

// GetOrders operation middleware
func (siw *ServerInterfaceWrapper) GetOrders(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.GetOrders(c)
}

// PostOrders operation middleware
func (siw *ServerInterfaceWrapper) PostOrders(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.PostOrders(c)
}

// DeleteOrdersId operation middleware
func (siw *ServerInterfaceWrapper) DeleteOrdersId(c *gin.Context) {

	var err error

	// ------------- Path parameter "id" -------------
	var id int32

	err = runtime.BindStyledParameter("simple", false, "id", c.Param("id"), &id)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter id: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.DeleteOrdersId(c, id)
}

// GetOrdersId operation middleware
func (siw *ServerInterfaceWrapper) GetOrdersId(c *gin.Context) {

	var err error

	// ------------- Path parameter "id" -------------
	var id int32

	err = runtime.BindStyledParameter("simple", false, "id", c.Param("id"), &id)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter id: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.GetOrdersId(c, id)
}

// PutOrdersId operation middleware
func (siw *ServerInterfaceWrapper) PutOrdersId(c *gin.Context) {

	var err error

	// ------------- Path parameter "id" -------------
	var id int32

	err = runtime.BindStyledParameter("simple", false, "id", c.Param("id"), &id)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter id: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.PutOrdersId(c, id)
}

// GinServerOptions provides options for the Gin server.
type GinServerOptions struct {
	BaseURL      string
	Middlewares  []MiddlewareFunc
	ErrorHandler func(*gin.Context, error, int)
}

// RegisterHandlers creates http.Handler with routing matching OpenAPI spec.
func RegisterHandlers(router gin.IRouter, si ServerInterface) {
	RegisterHandlersWithOptions(router, si, GinServerOptions{})
}

// RegisterHandlersWithOptions creates http.Handler with additional options
func RegisterHandlersWithOptions(router gin.IRouter, si ServerInterface, options GinServerOptions) {
	errorHandler := options.ErrorHandler
	if errorHandler == nil {
		errorHandler = func(c *gin.Context, err error, statusCode int) {
			c.JSON(statusCode, gin.H{"msg": err.Error()})
		}
	}

	wrapper := ServerInterfaceWrapper{
		Handler:            si,
		HandlerMiddlewares: options.Middlewares,
		ErrorHandler:       errorHandler,
	}

	router.GET(options.BaseURL+"/orders", wrapper.GetOrders)
	router.POST(options.BaseURL+"/orders", wrapper.PostOrders)
	router.DELETE(options.BaseURL+"/orders/:id", wrapper.DeleteOrdersId)
	router.GET(options.BaseURL+"/orders/:id", wrapper.GetOrdersId)
	router.PUT(options.BaseURL+"/orders/:id", wrapper.PutOrdersId)
}
