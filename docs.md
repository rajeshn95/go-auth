# How

### Commands i ran:

- `go mod init github.com/rajeshn95/access-in-go`

### Install the migrator

- `brew install golang-migrate`

### Install database

- go get github.com/jmoiron/sqlx
- go get github.com/jackc/pgx/v4/stdlib
- `make migrate.create name=create_books_table`
- `make make migrate.up`
- `make migrate.down`

### Install the fiber

express like server in go lang

- `go get -u github.com/gofiber/fiber/v3`
