# sample-go

Run app with

`go mod tidy`

`go run <path>/main.go`

Run tests with

`go mod tidy`

`go test <path>/sample-go/tests`

## Previous
Install dependency `golang-migrate` to create migrations. 
For create new migration:

`migrate create -ext sql -dir migrate/versions -seq <name for migration>`
