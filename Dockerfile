FROM golang:1.23.3 AS builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o main .

ENTRYPOINT ["./main"]