build:
	docker-compose build app

run:
	docker-compose up --build app

migrate:
	migrate -path ./migrations/pg -database 'postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable' up

drop:
	migrate -path ./migrations/pg -database 'postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable' down

test:
	go test ./...