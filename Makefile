run:
	docker-compose up --build app

test:
	go test -coverprofile=coverage.out ./...
	make test.coverage

test.coverage:
	go tool cover -html=coverage.out

swag:
	swag init -g internal/app/app.go

lint:
	golangci-lint run

mock:
	mockgen -source=internal/service/service.go -destination=internal/service/mocks/mock.go
	mockgen -source=internal/repository/repository.go -destination=internal/repository/mocks/mock.go