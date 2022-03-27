run:
	go run main.go 

postgres: 
	docker run --name crawler -p 8082:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=password -d postgres:14-alpine

createdb:
	docker exec -it crawler createdb --username=root --owner=root crawler

dropdb:
	docker exec -it crawler dropdb --username=root crawler

migrateup: 
	migrate -path db/migration -database "postgresql://root:password@localhost:8082/crawler?sslmode=disable" -verbose up

migratedown: 
	migrate -path db/migration -database "postgresql://root:password@localhost:8082/crawler?sslmode=disable" -verbose down

sqlc:
	docker run --rm -v C:\Users\amir\.vscode\bahar:/src -w /src kjconroy/sqlc generate

.PHONY: run	postgres createdb dropdb migrateup migratedown