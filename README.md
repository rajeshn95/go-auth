# How

### Commands to initialize ran:

- `go mod init github.com/rajeshn95/go-auth`

### Install the fiber

express like server in go lang

- `go get -u github.com/gofiber/fiber/v3`
- `go get github.com/gofiber/jwt/v2`
- `go get github.com/arsmn/fiber-swagger/v2`
- `go get github.com/golang-jwt/jwt`
- `go get github.com/go-playground/validator/v10`
- `go get github.com/joho/godotenv/autoload`

### makefile

Make file is used to store all the commands and commands stored can be ran in CLI wit prefix `make`:

- `make command`

### database

- Install the migrator: `brew install golang-migrate`
- Install database driver and sql:

  - `go get github.com/jmoiron/sqlx`
  - `go get github.com/jackc/pgx/v4/stdlib`

- `./platform/database` folder with database configuration (PostgreSQL)
- `./platform/database/migrations` folder with migration files (used with [golang-migrate/migrate](https://github.com/golang-migrate/migrate) tool)
- Migration commands:

  - `make migrate.create name=create_books_table`
  - `make migrate.up`
  - `make migrate.down`

### Server
