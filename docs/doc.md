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
* [x] refactor
* [x] basic api testing
* [x] database asseration function [ING] 
* [x] add lib/factory 
* [ ] basic lib testing
    * [x] assertion
        * [ ] success test case
        * [ ] failed and present error msg test case 
    * [ ] config
    * [ ] database initial
    * [x] factory
    * [ ] Worker test
    * [ ] EventListener test (async)
    * [x] lib/email test by human , success
* [x] Queue dispatcher (sync and async both finished)
* [x] double check assertion of "smartystreets/assertions" usage  
    and replace with "testify/assert"
* [ ] add command to handle job
  - [ ] Worker
    - [x] handle job
    - [x] worker command
    - [ ] handle error and retry
    - [x] add redis image
    - [x] add redis connection
    - [ ] add publisher function
  - [ ] Send to queue
  - [x] email sending services
* [ ] verify api
* [ ] ~~rabbitmq connection~~ 
* [ ] docker-compose.yml health check if db ready
* [ ] add grpc protocal
* [ ] Log file writing
  - [ ] usage of project `https://github.com/natefinch/lumberjack`
  - [ ] log level
* [ ] Finish async mailer flow
* [ ] Worker handler error job
* [ ] Refactoring
    * [ ] refactoring lib file structure  
        Ex: lib/ -> config/ -> config.go
    * [ ] refactor lib/asseration name to assertion
    * [ ] Refresh database in sequence
    * [ ] Add setting options for request read and write timeout
    * [ ] Add setting options for run mode
    * [ ] Standardize log format
    * [ ] add migration command
        - [ ] up
        - [ ] down
        - [ ] create
    
    
---
### Next Jobs
* [ ] Test EventListener - sync dispatch if triggered
* [ ] Redesign database schema
* [ ] Document flow and new schema
* [ ] Standardize log format
