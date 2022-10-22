postgres:
	docker run --name postgres14 -e  POSTGRES_USER=root -e  POSTGRES_PASSWORD=password -p 5433:5432 -d postgres:14.5-alpine

createdb:
	docker exec -it postgres14 createdb --username=root --owner=root tanzu

dropdb:
	docker exec -it postgres14 dropdb tanzu

deletedb:
	migrate -path db/migration -database "postgresql://root:password@localhost:5433/tanzu?sslmode=disable" -verbose force 1

migrateup:
	migrate -path db/migration -database "postgresql://root:password@localhost:5433/tanzu?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:password@localhost:5433/tanzu?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

.PHONY: postgres createdb dropdb migrateup migratedown sqlc test