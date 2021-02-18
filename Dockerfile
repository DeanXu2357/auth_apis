FROM golang:1.15

RUN go get -d -v ./...
RUN go install -v ./...

COPY . /go/src/app
WORKDIR /go/src/app

EXPOSE 8080

CMD ["./start.sh"]
