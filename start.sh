#!/bin/bash
go get -u github.com/swaggo/swag/cmd/swag
go mod tidy
go mod vendor
go build -o main .
./main serve