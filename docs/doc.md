# Design Blueprint

## Apis

* email register
* email verify
* email activate
* email resend
* email password recovery
* email password reset

## Schema

* users

|name|type|length|default|index|comment|
|---:|---:|---:|---:|---:|---:|
|id|string|32|N|PK|uuid v4 generate by `{root}/models/users`|
|email|string|128|N|UNIQUE||
|name|string|50|''|||
|created_at|timestamp|||||
|updated_at|timestamp|||||

* email_login

|name|type|length|default|index|comment|
|---:|---:|---:|---:|---:|---:|
|email|string|128|N|PK||
|pwd|string|255|N||hash, nullable|
|verifed_at|timestamp|||||
|created_at|timestamp|||||
|updated_at|timestamp|||||

* email_verify

|name|type|length|default|index|comment|
|---:|---:|---:|---:|---:|---:|
|id|int||N|PK|AutoIncreament|
|email|string|128|N|index||
|verification|smallint||0|N||0:未驗證, 1:已驗證|
|created_at|timestamp|||||
|updated_at|timestamp|||||

## todo list
[ ] docker-compose.yml health check if db ready
[ ] add grpc protocal
[ ] add new command to handle job
[x] refactor
[x] basic api testing
[ ] basic lib testing
[ ] verify api