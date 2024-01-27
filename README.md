<!-- db -->
1. 
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

2.
migrate create -ext sql -dir server/migrations -seq Product
migrate create -ext sql -dir server/migrations -seq Customer
migrate create -ext sql -dir server/migrations -seq Order

3.
migrate -database "postgres://postgres:postgres@localhost:5432/sabio-ekuator?sslmode=disable" -path server/migrations up

<!-- proto -->
protoc --go_out=. --go-grpc_out=. proto/*.proto
go build -o bin/server ./server
go build -o bin/client ./client

