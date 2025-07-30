# Wallet Management API - Requirement

## Overview

API for managing digital wallets (Wallets) that allows users to create, view, edit and delete wallets, supporting data on balance, currency and wallet status.

---

## Entities

### Wallet

| Field     | Type    | Description                       |
|-----------|---------|---------------------------------|
| Field     | Type    | Description                       |
|-----------|---------|---------------------------------|
| ID        | integer | Wallet ID (Primary Key)         |
| OwnerID   | integer | Owner's ID                      |
| Balance   | number  | Current balance in the wallet   |
| Currency  | string  | Currency type (e.g., THB, USD)  |
| Status    | string  | Wallet status (Active, Inactive)|
| CreatedAt | string  | Creation timestamp              |
| UpdatedAt | string  | Last update timestamp           |

---

## API Endpoints

### 1. Create Wallet

- **URL:** `/api/v1/wallets`
- **Method:** POST
- **Request Body:**

```json
{
  "currency": "THB",
  "owner_id": 123
}
```

- **Acceptance Criteria:**
  - `owner_id` (integer) and `currency` (string) are required.
  - `balance` is initialized to 0 and `status` to `Active`.
- **Response:**

```json
{
  "id": 1,
  "owner_id": 123,
  "balance": 0,
  "currency": "THB",
  "status": "Active",
  "created_at": "2025-07-30T10:00:00Z",
  "updated_at": "2025-07-30T10:00:00Z"
}
```

---

### 2. Get Wallet

- **URL:** `/api/v1/wallets/{owner_id}`
- **Method:** GET
- **Acceptance Criteria:**
  - Retrieves wallet information by `owner_id`.
  - Returns HTTP 404 if the wallet is not found.

- **Response:**

```json
{
  "id": 1,
  "owner_id": 123,
  "balance": 1000.50,
  "currency": "THB",
  "status": "Active",
  "created_at": "2025-07-30T10:00:00Z",
  "updated_at": "2025-07-30T10:00:00Z"
}
```

---

### 3. Update Wallet

- **URL:** `/api/v1/wallets/{owner_id}`
- **Method:** PUT
- **Request Body:**

```json
{
  "balance": 1500.75,
  "currency": "USD",
  "status": "Inactive"
}
```

- **Acceptance Criteria:**
  - Allows updating `balance` (number), `currency` (string), and `status` (string).
  - `currency` must be a supported currency.
  - `status` must be one of `Active` or `Inactive`.
  - Updates the wallet information and returns the updated wallet.
  - Returns HTTP 404 if the wallet is not found.

- **Response:**

```json
{
  "id": 1,
  "owner_id": 123,
  "balance": 1000.50,
  "currency": "USD",
  "status": "Inactive",
  "created_at": "2025-07-30T10:00:00Z",
  "updated_at": "2025-08-01T08:00:00Z"
}
```

---

### 4. Delete Wallet

- **URL:** `/api/v1/wallets/{owner_id}`
- **Method:** DELETE
- **Acceptance Criteria:**
  - Deletes or deactivates the wallet by `owner_id`.
  - Actual deletion might be a soft-delete or a status change to `Inactive`.
  - Returns HTTP 404 if the wallet is not found.
  - Returns HTTP 204 No Content on successful deletion.

---

## Technology Stack

- **Backend Framework:** Gin (HTTP Router)  
- **ORM:** GORM (Go ORM)  
- **Architecture:** Hexagonal Architecture (Clean Architecture)


## Project Structure (Conceptual)

This project follows a Hexagonal Architecture (also known as Ports and Adapters or Clean Architecture) to ensure a clear separation of concerns, testability, and maintainability.

```
.
├── cmd/                     # Application Entry Points
├── config/                  # Configuration Management
├── docs/                    # API Documentation
├── infrastructure/          # Infrastructure Layer (Driven Adapters)
│   # Concrete implementations of ports (e.g., database, external services)
├── internal/                # Internal Application Logic
│   ├── adapters/            # Driving and Driven Adapters
│   │   ├── http/            # Driving Adapter (HTTP API)
│   │   └── repository/      # Driven Adapter (Database Repository)
│   ├── core/                # Core Business Logic (Hexagon)
│   │   ├── domain/          # Domain Entities and Business Rules
│   │   ├── port/            # Ports (Interfaces for Core to interact with Adapters)
│   │   │   ├── repository.go # Interface for data persistence (Driven Port)
│   │   │   └── service.go   # Interface for business logic (Driving Port)
│   │   └── service/         # Application Services (Use Cases)
│   └── dto/                 # Data Transfer Objects
├── pkg/                     # Reusable Packages
├── protocol/                # Protocol-Specific Implementations (e.g., HTTP server setup)
└── utils/                   # General Utilities
```

---

## Getting Started

Follow these steps to set up and run the Wallet Management API locally:

1.  **Clone the Repository:**
    ```bash
    git clone https://github.com/your-username/go-api-crud.git
    cd go-api-crud
    ```

2.  **Run with Docker Compose:**
    This will build the Docker images, set up the database, and start the API server.
    ```bash
    docker-compose up --build
    ```

3.  **Access API Documentation:**
    Once the application is running, open your browser and navigate to the Swagger UI:
    `http://localhost:8080/swagger/index.html`

4. **Access DB Using UI:**
    Open your database tool and using url below to connect with db
    `postgresql://user_wallet_user:user_wallet_pass@localhost/wallets?statusColor=&env=local&name=wallet&tLSMode=0&usePrivateKey=false&safeModeLevel=0&advancedSafeModeLevel=0&driverVersion=0&showSystemSchemas=0&driverVersion=0&lazyload=False`
---

## Running Tests

To run all unit tests, use the following command:

```bash
go test -v -cover ./...
```

To generate mocks (required if interfaces change):

```bash
go install github.com/vektra/mockery/v2@latest
go generate ./...
```

---