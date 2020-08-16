# Api 練習 -- 登入微服務
## 啟動專案
* 初次啟動  
`go build -o main . && sudo docker-compose up -d`
* 變更後重新編譯  
`sudo docker-compose exec app go build -o main && sudo docker-compose restart  app`
* Run test  
`go test ./...`

## About Migrate  
`https://github.com/golang-migrate/migrate/tree/master/cmd/migrate`
* migrate from local 
`./cmd/migrate.linux-amd64 -database "postgres://postgres:fortestpwd@localhost:45487/auth?sslmode=disable" -verbose -path db/migrations up`

## Q&A

* migration failed (fix and force version)  
run  `drop table schema_migrations;` in db
