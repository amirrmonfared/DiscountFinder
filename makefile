run:
	go run ./cmd/web/.
postgres: 
	docker run --name crawler -p 8082:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=password -d postgres:14-alpine

createdb:
	docker exec -it crawler createdb --username=root --owner=root crawler

dropdb:
	docker exec -it crawler dropdb --username=root crawler

psql:
	docker exec -it crawler psql -U root -d crawler

migrateup: 
	migrate -path db/migration -database "postgresql://root:password@localhost:8082/crawler?sslmode=disable" -verbose up

migratedown: 
	migrate -path db/migration -database "postgresql://root:password@localhost:8082/crawler?sslmode=disable" -verbose down

sqlc:
	docker run --rm -v C:\Users\amir\.vscode\projects\bahar:/src -w /src kjconroy/sqlc generate

test:
	go test -v -cover ./...

mock:
	mockgen -destination db/mock/store.go github.com/amirrmonfared/DiscountFinder/db/sqlc Store

root:
	docker exec -it crawler /bin/sh

db_schema: 
	migrate create -ext sql -dir db/migration -seq init_schema

.PHONY: run	postgres createdb dropdb migrateup migratedown sqlc mock test root psql db_schema psql