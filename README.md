# Api practice -- login service
## Start project
1. initial  
`sudo docker-compose up -d`
2. do migration  
`sudo docker-compose exec app go run main.go migrate up`
3. copy config  
`cp config.yml.example config.yml`

* refresh code   
`sudo docker-compose restart  app`
* Run test (since integration tests, refresh db not set parallel 1 will crash)  
`go test -p 1 ./...`

## Tools
* Jaeger UI  
`http://localhost:16686`
* Swagger 
`http://localhost:666/swagger/index.html`

## Commands
* `./main serve`
* `./main work:email`
* `./main migrate up`
* `./main migrate refresh`
* `./main migrate down [steps, default all]`
* `./main migrate make [make params]`

## About [Migrate](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate)
* migrate from local   
`./cmd/migrate.linux-amd64 -database "postgres://postgres:fortestpwd@localhost:45487/auth?sslmode=disable" -verbose -path db/migrations up`
* migrate in container  
`sudo docker-compose exec app go run main.go migrate up`

## generate key and secret 
1. Generate private key  
`ssh-keygen -t rsa -b 4096 -m PEM -f jwtRS256.key`  
2. Generate public key  
`openssl rsa -in jwtRS256.key -pubout -outform PEM -out jwtRS256.key.pub`  
3. Setting  
`copy and paste to config.yml`

## Q&A

* migration failed (fix and force version)  
run  `drop table schema_migrations;` in db
