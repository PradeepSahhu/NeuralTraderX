# NeuralTraderX

Purpose of this Application?

This application will be a multi agent trading system that will use the LLMs to make the decision of buying or selling stocks mainly in intraday trading.

Which LLMs will be used?

Will make it to support all llms of every big organization like google's gemini, openai's gpt, meta's llama, and more. but for the development phase will use the deepseek's flash model. (cost effective)

Language?

The application will be build using golang (purely backend) and will be no UI

any broker api?
yes, will be using the upstocks api for the trading part.

How will the system work?
TBD...

currently only exploring the upstocks api endpoints and figuring out how to use them
will also be having a bruno collection or a http file with the api endpoints of this application and the upstocks api endpoints for easy reference and testing

Any Coding Agent will be used?
NO...

### Technologies

- Database PostgreSQL - ORM = GORM
- LLMs = Deepseek's Flash (for development phase)
- Broker API = Upstocks API / Angel One
- Programming Language = Golang
- Containerization = Docker

#### Database setup

Install the golang-migration tool for database migrations

```bash
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
```

install the ORM (GORM)

```bash
go get -u gorm.io/gorm
go get -u gorm.io/driver/postgres
```

#### Create the migration scripts

```bash
migrate create -ext sql -dir migrations <name>

```

#### Run the migration scripts

```bash
migrate -path ./migrations -database "postgres://pradeep:password@localhost:5432/stocks?sslmode=disable" up
```
