FROM golang:1.14

WORKDIR /go/src/app
COPY . .

#RUN go get -d -v ./...
#RUN go install -v ./...

EXPOSE 8080

#ENTRYPOINT ["go", "build", "-mod", "vendor", "-o", "main"]
RUN go build -mod vendor -o main

CMD ["./main"]
#ENTRYPOINT ["./main"]
