# An API Server boilerplate for Golang

Authored by [Marcus Ong](https://github.com/mraacus)

A production-ready Go API server boilerplate using Echo, sqlc, and Goose.

## Technologies

This repository uses the following technology stack.

- **[Echo](https://echo.labstack.com/)**: High performance, minimalist Go web framework
- **[sqlc](https://docs.sqlc.dev/)**: Generate type-safe Go code from SQL
- **[Goose](https://github.com/pressly/goose)**: Database migration tool
- **[PostgreSQL](https://www.postgresql.org/)**: Robust, open-source relational database

## Prerequisites

- Go 1.21+
- Docker
- Goose + sqlc
- Make (optional, for using Makefile commands)

## Getting Started

### Environment Setup

1. Clone the repository

   ```bash
   git clone https://github.com/mraacus/go-api-boilerplate.git
   ```

2. Create your .env file
   ```bash
   cp .env.sample .env
   # Edit .env with your preferred settings
   ```

### Set up your PostgreSQL database with Docker

This boilerplate comes with a docker-compose file to spin up a local PostgreSQL database.

1. Start the PostgreSQL container

   ```bash
   make docker-up
   # or
   docker-compose up -d
   ```

2. Run database migrations
   ```bash
   goose up
   ```

### Running the server

For development with hot reloading:
`bash
    make watch
    `
Or run the server directly:
`bash
    go run cmd/api/main.go
    `

### Test the server API

Testing just the API connection:

````
GET /

    Response:
    {
    "message": "I am groot"
    }
    ```

Testing database functionality:

````

POST /users

    Request Body:
    {
    "name": "username",
    "role": "admin"
    }

    Response:
    {
    "id": 1,
    "name": "username",
    "role": "admin"
    }

```

## Project Structure

```

.
├── cmd/
│ └── api/ # Application entrypoint
├── internal/
│ ├── database/ # Database connection and queries
│ │ ├── migrations/ # Database migrations
│ │ ├── queries/ # sqlc generated go code
│ │ └── sqlc/ # sqlc queries and sqlc.yaml
│ ├── modules/ # Modules Domains
│ └── server/ # Server setup
├── .env # Environment variables
├── docker-compose.yml # Docker services
└── Makefile # Development Make commands

```

```
