# Checkout Service

A Go microservice that implements a supermarket checkout:

> Scan a sequence of item SKUs (A, B, C, D) and calculate the total price,  
> applying special “n for y” pricing rules (e.g. 3 × A for 130 pence).

## Features

- Implements `service.ICheckout` (`Scan` + `GetTotalPrice`)  
- Single HTTP POST `/checkout` endpoint accepts JSON array of SKUs  
- Validates each SKU (uppercase, single‑char) via Gin binding + validator  
- Returns `200 {"total":…}` on success, `400/500 {"error":…}` on failure  
- Auto‑generated Swagger UI at `/docs`

## Prerequisites

- Go 1.24+  
- [swag](https://github.com/swaggo/swag) CLI to regenerate docs

## Installation

git clone https://github.com/AdiaElena/checkout-service.git
cd checkout-service
go mod tidy

## Running the service

go run cmd/api/main.go

## Regenerate swagger docs

swag init -g cmd/api/main.go -o docs

## Links

- Swagger UI: http://localhost:8080/docs/index.html

## Example call

- API: POST http://localhost:8080/checkout
    - Header: Content-Type: application/json
    - Body: {"skus":["A","B","B","C"]}
    - Result: {"total":115}

## Development and testing
Before you commit, you can run the following checks from your project root (where go.mod lives):
# 1. Format & imports
go fmt ./...

# 2. Vet & lint
go vet ./...

# 3. Module hygiene
go mod tidy

# 4. Regenerate Swagger (if any handler/DTOs changed)
swag init -g cmd/api/main.go -o docs

# 5. Tests
go test ./... -v -race -timeout 30s
