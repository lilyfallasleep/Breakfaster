FROM golang:1.14.8-alpine3.11 AS builder

RUN apk update && apk add alpine-sdk git && rm -rf /var/cache/apk/*
RUN mkdir -p /app
WORKDIR /app
COPY Makefile .
RUN make swag

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN make doc build

FROM alpine:latest

RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*

RUN mkdir -p /app
WORKDIR /app
COPY --from=builder /app/cmd/server/breakfaster /app/script/entrypoint.sh /app/script/wait-for-it.sh ./
RUN apk add --no-cache bash
RUN chmod +x wait-for-it.sh entrypoint.sh

ENTRYPOINT ["./wait-for-it.sh", "db:3306", "--", "./wait-for-it.sh", "redis-node1:7000", "--", "./entrypoint.sh"]