# Configuration Guide

## 📁 Config Files Structure

```
├── config.example.json    # Template for team (committed)
├── config.json           # Production config (NOT committed)
├── config.dev.json       # Development config (can be committed)
├── config.test.json      # Testing config (committed)
└── config.prod.json      # Production config (NOT committed)
```

## Quick Setup

### For New Developers:

```bash
# Linux/Mac
./setup.sh

# Windows
setup.bat

# Or manually
cp config.example.json config.json
# Edit config.json with your actual values
```

## Environment Usage

### Local Development

```bash
# Option 1: Use config.dev.json (recommended)
export APP_ENV=development
go run cmd/web/main.go

# Option 2: Use environment variables
export APP_ENV=development
export PAKYUS_DB_USERNAME=axata
export PAKYUS_DB_PASSWORD=yourpassword
export PAKYUS_DB_NAME=pakyus_commerce_dev
go run cmd/web/main.go
```

### Testing

```bash
export APP_ENV=testing
go test ./test/...
```

### Production

```bash
# Uses config.json by default
go run cmd/web/main.go

# Or with env override
export PAKYUS_DB_PASSWORD=prod_password
export PAKYUS_DB_HOST=prod.database.com
go run cmd/web/main.go
```

## Environment Variables

All config values can be overridden with environment variables:

| Config Path               | Environment Variable   | Example          |
| ------------------------- | ---------------------- | ---------------- |
| `database.username`       | `PAKYUS_DB_USERNAME`   | `postgres`       |
| `database.password`       | `PAKYUS_DB_PASSWORD`   | `secretpass`     |
| `database.host`           | `PAKYUS_DB_HOST`       | `localhost`      |
| `database.port`           | `PAKYUS_DB_PORT`       | `5432`           |
| `database.name`           | `PAKYUS_DB_NAME`       | `mydb`           |
| `web.port`                | `PAKYUS_WEB_PORT`      | `3000`           |
| `kafka.bootstrap.servers` | `PAKYUS_KAFKA_SERVERS` | `localhost:9092` |

## CI/CD Setup

### GitHub Actions Example:

```yaml
name: CI
on: [push, pull_request]

jobs:
  test:
    runs-on: ubuntu-latest
    services:
      postgres:
        image: postgres:15
        env:
          POSTGRES_PASSWORD: postgres
          POSTGRES_DB: pakyus_commerce_test
        options: >-
          --health-cmd pg_isready
          --health-interval 10s
          --health-timeout 5s
          --health-retries 5
        ports:
          - 5432:5432

    steps:
      - uses: actions/checkout@v3
      - uses: actions/setup-go@v4
        with:
          go-version: "1.21"

      - name: Run tests
        env:
          APP_ENV: testing
          PAKYUS_DB_USERNAME: postgres
          PAKYUS_DB_PASSWORD: postgres
          PAKYUS_DB_HOST: localhost
          PAKYUS_DB_PORT: 5432
          PAKYUS_DB_NAME: pakyus_commerce_test
        run: |
          go test ./test/...
```

## Security Best Practices

### DO:

- Commit `config.example.json`, `config.dev.json`, `config.test.json`
- Use environment variables for sensitive data in production
- Use different database names for different environments
- Document all required environment variables

### DON'T:

- Commit `config.json` (production config)
- Put sensitive data in committed config files
- Use production database for development/testing
- Hardcode passwords or API keys

## Migration Commands

```bash
# Development
APP_ENV=development migrate -path db/migrations -database "postgres://axata:password@localhost:5433/pakyus_commerce_dev?sslmode=disable" up

# Testing
APP_ENV=testing migrate -path db/migrations -database "postgres://postgres:postgres@localhost:5432/pakyus_commerce_test?sslmode=disable" up

# Production (use env vars)
migrate -path db/migrations -database "postgres://$PAKYUS_DB_USERNAME:$PAKYUS_DB_PASSWORD@$PAKYUS_DB_HOST:$PAKYUS_DB_PORT/$PAKYUS_DB_NAME?sslmode=disable" up
```
