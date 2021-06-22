# sample-go

Directory app: REST Api with PostgreDB connection and AMQP publisher
Directory consumer: AMQP Consumer with PostgreDB connection

Run **app** with

`go mod tidy`

`go run <path>/app/main.go`

tests `go test <path>/sample-go/app/tests`


Run **consumer** with

`go mod tidy`

`go run <path>/consumer/main.go`

tests `go test <path>/sample-go/consumer/tests`

## Migration
To create new migration file, install dependency `golang-migrate` and execute:

`migrate create -ext sql -dir migrate/versions -seq <name for migration>`

## AMQP
Ref. https://dev.to/koddr/working-with-rabbitmq-in-golang-by-examples-2dcn#setting-up-a-message-consumer