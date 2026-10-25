# Saffron Real Estate

An in-progress Go backend for managing real estate purchase requests and users, backed by PostgreSQL.

The current entry point connects to PostgreSQL, creates the application database if needed, applies migrations, and exits. HTTP routes and a frontend are not implemented yet.

## Stack

- Go 1.25.0
- PostgreSQL 18
- `pgx` for database access
- `golang-migrate` for SQL migrations
- Docker Compose for local services

## Run with Docker

Install Docker with Docker Compose, then run from the repository root:

```sh
docker compose up --build
```

Compose starts PostgreSQL, waits for it to become healthy, and runs the backend. The backend creates `saffron_backend_db` and applies the migrations in `backend/migrations`. A successful backend run exits with code `0`; PostgreSQL continues running.

Although Compose maps port `8080`, the backend does not currently start an HTTP server.

To connect to the application database:

```sh
docker compose exec db psql -U app -d saffron_backend_db
```

Stop the services with:

```sh
docker compose down
```

Database data persists in the `postgres_data` named volume across restarts.

## Run locally

Install Go compatible with `backend/go.mod` and run a PostgreSQL 18 instance accessible from your machine.

Create `backend/.env` with a connection string for your local PostgreSQL instance:

```dotenv
DATABASE_URL=postgres://app:devpassword@localhost:5432/app?sslmode=disable
```

The example credentials match Compose, but Compose does not publish PostgreSQL's port to the host. Running the Go process locally requires a separately accessible PostgreSQL instance or a local Compose override that publishes the database port.

From the repository root:

```sh
cd backend
go mod download
go run ./cmd/api
```

Run from `backend/` so the process can find `.env` and the relative `migrations/` directory. You can also supply `DATABASE_URL` through the environment; existing environment variables take precedence over `.env` values.

## Project structure

```text
.
├── compose.yaml
└── backend/
    ├── cmd/api/             # Database initialization entry point
    ├── internal/
    │   ├── config/          # Environment configuration
    │   ├── db/              # Database creation, pools, and migrations
    │   └── model/           # User and purchase request models
    ├── migrations/          # Ordered SQL migrations
    ├── Dockerfile
    ├── go.mod
    └── go.sum
```
