# Cheki Service

A clean architecture Golang REST API service.

## Architecture

This project follows Clean Architecture principles with the following structure:

```
cheki-service/
├── cmd/api/                    # Application entry points
├── internal/                   # Private application code
│   ├── domain/                 # Business logic layer
│   │   ├── entity/            # Business entities
│   │   └── repository/        # Repository interfaces
│   ├── usecase/               # Application business rules
│   ├── delivery/              # Presentation layer
│   │   └── http/              # HTTP handlers
│   └── infrastructure/        # External concerns
│       ├── database/          # Database implementations
│       └── config/            # Configuration
├── pkg/                       # Public libraries
│   ├── logger/               # Logging utilities
│   └── utils/                # Common utilities
├── migrations/               # Database migrations
├── docs/                    # Documentation
└── test/                    # Test files
```

## Layers

### Domain Layer (`internal/domain/`)
- **Entities**: Core business objects
- **Repositories**: Interfaces for data access

### Use Case Layer (`internal/usecase/`)
- Application-specific business rules
- Orchestrates data flow between layers

### Delivery Layer (`internal/delivery/`)
- HTTP handlers and routes
- Request/response formatting

### Infrastructure Layer (`internal/infrastructure/`)
- Database connections and implementations
- External services
- Configuration management

## Getting Started

### Prerequisites
- Go 1.24+
- PostgreSQL (or your preferred database)

### Installation

1. Clone the repository
2. Install dependencies:
```bash
go mod tidy
```

3. Set environment variables:
```bash
export DB_HOST=localhost
export DB_PORT=5432
export DB_USER=postgres
export DB_PASSWORD=password
export DB_NAME=chekisvc
export SERVER_PORT=8080
```

4. Run the application:
```bash
go run cmd/api/main.go
```

## API Endpoints

- `GET /api/v1/health` - Health check
- `POST /api/v1/users` - Create user
- `GET /api/v1/users/:id` - Get user by ID
- `PUT /api/v1/users/:id` - Update user
- `DELETE /api/v1/users/:id` - Delete user
- `GET /api/v1/users` - List users (with pagination)

## Dependencies

Add the following dependencies to your `go.mod`:

```go
require (
    github.com/gin-gonic/gin v1.9.1
    gorm.io/gorm v1.25.0
    gorm.io/driver/postgres v1.5.0
    golang.org/x/crypto v0.0.0-20230315142452-642cacee5cc0
)
```

## Testing

Run tests with:
```bash
go test ./...
```

## Contributing

1. Fork the repository
2. Create a feature branch
3. Commit your changes
4. Push to the branch
5. Create a Pull Request
