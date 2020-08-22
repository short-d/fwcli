FROM golang:1.13.1-alpine AS builder

WORKDIR /approot

RUN apk add --no-cache git bash

# Install dependencies
COPY go.mod go.sum ./
RUN go mod download

# Verify dependencies
RUN go mod verify

COPY . .

RUN go build -o build/app main.go

FROM alpine:3.10

WORKDIR /approot

RUN apk add --no-cache bash

COPY --from=builder /approot/build/app ./build/app
COPY --from=builder /approot/app/adapter/sqldb/migration ./app/adapter/sqldb/migration
COPY --from=builder /approot/app/adapter/gqlapi/schema.graphql ./app/adapter/gqlapi/schema.graphql