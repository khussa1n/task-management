FROM golang:1.20-buster

RUN go version
ENV GOPATH=/

COPY ./ ./

RUN apt-get update
RUN apt-get -y install postgresql-client

RUN go mod download
RUN go build -o task-management ./cmd/task-management/main.go

CMD ["./task-management"]