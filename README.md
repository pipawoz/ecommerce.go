# E-Commerce Order Processing System - Golang

A scalable order processing system built using Golang, Temporal for microservice orchestration, Gin for the web framework, Postgres for the database, and Docker for containerization.

## Table of Contents
- [Introduction](#introduction)
- [Features](#features)
- [Installation](#installation)
- [Usage](#usage)
- [Dependencies](#dependencies)
- [Configuration](#configuration)
- [License](#license)

## Introduction
This project implements a robust and scalable order processing system in Golang. It supports CRUD operations for orders, leverages Temporal for orchestrating workflows, and utilizes Postgres as the database.

## Features
- RESTful CRUD API using OpenAPI and Gin.
- Microservice orchestration with Temporal.
- Database migrations and SQL operations using sqlc.
- Docker setup for easy environment management.
  
## Installation
1. Clone the repository:
   ```bash
   git clone https://github.com/pipawoz/order-processing-system
   cd order-processing-system
   ```

2. Install dependencies:
   ```bash
   go mod download
   ```

3. Build and run Docker containers:
   ```bash
   docker-compose up
   ```

## Usage
- Run the application:
  ```bash
  go run cmd/api/main.go
  ```

- Use `curl` or Postman to interact with the API:
  - Create an order: 
    ```bash
    curl -X POST http://localhost:8080/orders -H "Content-Type: application/json" -d '{"customer_id": 1, "total_amount": 100.50}'
    ```

- The Temporal workflow starts automatically after order creation.

## Dependencies
- [Golang](https://golang.org/)
- [Gin](https://github.com/gin-gonic/gin)
- [Temporal](https://temporal.io/)
- [PostgreSQL](https://www.postgresql.org/)
- [Docker](https://www.docker.com/)
- [sqlc](https://github.com/kyleconroy/sqlc)
- [oapi-codegen](https://github.com/deepmap/oapi-codegen)

## Configuration
- The environment variables for PostgreSQL and Temporal are managed in `docker-compose.yml`. For local development, Postgres runs on `localhost:5432` and Temporal on `localhost:7233`.

## License
This project is licensed under the MIT License. See the LICENSE file for details.
