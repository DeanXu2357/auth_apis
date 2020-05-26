FROM golang:1.14

# RUN mkdir /app
WORKDIR /app
COPY . .

RUN go mod tidy
RUN go mod download

RUN go build -o main .

EXPOSE 8080

ENTRYPOINT ["./main"]

CMD ["./main"]