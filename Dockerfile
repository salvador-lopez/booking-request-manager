FROM golang:1.17-alpine

WORKDIR /go/src/booking-request-manager/cmd/api

CMD go run main.go
