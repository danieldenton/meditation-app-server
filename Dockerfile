# syntax=docker/dockerfile:1

FROM golang:1.23-alpine

WORKDIR /app

RUN go install github.com/air-verse/air@latest

COPY . .

RUN go mod tidy

EXPOSE 8080

CMD ["./main"]