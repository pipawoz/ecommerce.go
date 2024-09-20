package main

import (
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/pipawoz/ecommerce.go/internal/service"
)

func main() {
	svc, err := service.NewService()
	if err != nil {
		log.Fatalf("Failed to initialize service: %v", err)
	}
	defer svc.Close()

	r := gin.Default()
	svc.Handler.RegisterRoutes(r)

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}