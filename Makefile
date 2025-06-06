postgres:
	docker run --name linh-postgres -e POSTGRES_PASSWORD=secret -e POSTGRES_USER=root -p 5432:5432 -d postgres

postgres-start:
	docker start linh-postgres

createdb:
	docker exec -it linh-postgres createdb --username=root --owner=root simple_bank_go

dropdb:
	docker exec -it linh-postgres dropdb simple_bank_go

migrateup:
	migrate -path db/migrations -database "postgresql://root:secret@localhost:5432/simple_bank_go?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migrations -database "postgresql://root:secret@localhost:5432/simple_bank_go?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:	
	go test -v -coverpkg=./... ./...

.PHONY: postgres createdb dropdb migrateup migratedown sqlc