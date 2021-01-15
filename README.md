# Api practice -- login service
## Start project
* initial  
`sudo docker-compose up -d`
* do migration  
`sudo docker-compose exec app go run main.go migrate up`
* refresh dev environment 
`sudo docker-compose exec app go build -o main && sudo docker-compose restart  app`
* Run test (since integration tests, refresh db not set parallel 1 will crash)  
`go test -p 1 ./...`

## Commands
* `./main serve`
* `./main work:email`
* `./main migrate up`
* `./main migrate refresh`
* `./main migrate down [steps, default all]` 
* `./main migrate make [make params]`

## About Migrate  
`https://github.com/golang-migrate/migrate/tree/master/cmd/migrate`
* migrate from local   
`./cmd/migrate.linux-amd64 -database "postgres://postgres:fortestpwd@localhost:45487/auth?sslmode=disable" -verbose -path db/migrations up`

## generate key and secret 
`ssh-keygen -t rsa -b 4096 -m PEM -f jwtRS256.key`  
`openssl rsa -in jwtRS256.key -pubout -outform PEM -out jwtRS256.key.pub`  
`copy and paste to config.yml`

## Q&A

* migration failed (fix and force version)  
run  `drop table schema_migrations;` in db
