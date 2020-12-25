FROM golang:1.15

RUN go get -d -v ./...
RUN go install -v ./...

COPY . /go/src/app
WORKDIR /go/src/app

#RUN go build -o main .

EXPOSE 8080

CMD ["./start.sh"]
