postgres:
	docker run --name postgres777 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASS=root -d postgres

createdb:
	docker exec -it postgres777 createdb --username=root --owner=root bank

dropdb:
	docker exec -it postgres777 dropdb bank

migrateup:
	migrate -path db/migration -database "postgresql://root:root@localhost:5432/bank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:root@localhost:5432/bank?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

# make unit_test TEST=TestFunctionName
unit_test:
	go test -v -cover ./db/sqlc -run $(TEST)

server:
	go run main.go

mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/narymbaev/techschool/db/sqlc Store

.PHONY: postgres createdb dropdb migrateup migratedown test sqlc server