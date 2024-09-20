package service

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/pipawoz/ecommerce.go/internal/api"
	"github.com/pipawoz/ecommerce.go/internal/db"
	"go.temporal.io/sdk/client"
)

type Service struct {
	DB            *sql.DB
	Queries       *db.Queries
	TemporalClient client.Client
	Handler       *api.Handler
}

func NewService() (*Service, error) {
	// Initialize database connection
	dbConn, err := sql.Open("postgres", "postgresql://orderuser:orderpass@postgres:5432/orderdb?sslmode=disable")
	if err != nil {
		return nil, err
	}

	temporalHost := os.Getenv("TEMPORAL_HOST")
	if temporalHost == "" {
		temporalHost = "temporal:7233"
	}

	fmt.Println("Connecting to Temporal at", temporalHost)

	// Initialize Temporal client
	temporalClient, err := client.Dial(client.Options{
		HostPort: temporalHost,
	})
	if err != nil {
		return nil, err
	}

	// Initialize queries
	queries := db.New(dbConn)

	// Initialize handler
	handler := api.NewHandler(queries, temporalClient)

	return &Service{
		DB:            dbConn,
		Queries:       queries,
		TemporalClient: temporalClient,
		Handler:       handler,
	}, nil
}

func (s *Service) Close() {
	s.DB.Close()
	s.TemporalClient.Close()
}