# Build stage
FROM golang:1.23.1-alpine AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /ecommerce.go ./cmd/api

# Run stage
FROM alpine:latest

WORKDIR /

COPY --from=build /ecommerce.go /ecommerce.go

EXPOSE 8080

# COPY wait-for-it.sh /usr/bin/wait-for-it
# RUN chmod +x /usr/bin/wait-for-it

# ENTRYPOINT ["/usr/bin/wait-for-it", "temporal:7233", "--", "/ecommerce.go"]

ENTRYPOINT ["/ecommerce.go"]