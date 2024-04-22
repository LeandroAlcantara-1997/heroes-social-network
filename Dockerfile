FROM golang:1.22-alpine as builder

RUN apk update
WORKDIR /src
COPY . .
COPY go.mod go.sum ./
RUN GOOS=linux go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o api main.go

FROM alpine:3.19 as run
WORKDIR /src
RUN apk update && apk add bash
COPY --from=builder /src/api ./
COPY --from=builder /src/config ./config




