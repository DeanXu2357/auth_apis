#!/bin/bash
go get -u github.com/swaggo/swag/cmd/swag
go build -o main .
./main serve