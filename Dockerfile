FROM golang:1.20.6-alpine3.17 AS builder
WORKDIR /app
COPY . .
RUN go build -o task-management ./cmd/task-management/main.go


FROM alpine:3.14
WORKDIR /app
COPY --from=builder /app/task-management .
COPY config.yaml .
COPY migrations/pg/000001_init.up.sql /app/migrations/pg/000001_init.up.sql
COPY migrations/pg/000001_init.down.sql /app/migrations/pg/000001_init.down.sql
EXPOSE 8080
RUN apk update && apk add postgresql-client
CMD ["sh", "-c", "while ! nc -z db 5432; do sleep 1; done && ./task-management"]