# Design Blueprint

## Apis

* email register
* email verify
* email activate
* ~~email resend~~
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
|id|uuid||N|PK|uuid|
|email|string|128|N|||
|mail_type|string|64|N||verify, reset|
|verification|smallint||0|N||0:未驗證, 1:已驗證|
|created_at|timestamp|||||
|updated_at|timestamp|||||

## flow
* Register by email
    1. request api `/api/v1/email/register`
    2. model user create
    3. model email_login create
    4. model email_verify create
    5. send email
        - url = `/api/v1/email/activate?token=` + token(使用 email_verify id as id to generate jwt token)
    6. user request verify url`/api/v1/email/activate` 
    7. parse token
    8. query email_verify by id , and verify
        1. email if the same 
        2. created_at compare now if in setting durations
        3. mail_type is `verify`
    9. update email_login
    10. delete email_verify
    11. redirect to login page (setting in config file)
* Reset password
    1. request recovery password api `/api/v1/email/recovery`
    2. model email_verify create 
    3. send email
        - url = `/api/v1/email/activate?token=` + token(使用 email_verify id as id to generate jwt token)    
    4. request reset password api `/api/v1/email/reset`
    5. query email_verify by id
    6. verify
        1. email
        2. type
        3. created_at
    7. update email_login
    8. delete email_verify
    9. redirect to login page (setting in config file)

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
    * [ ] Test EventListener - sync dispatch if triggered
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
* [ ] Log file writing
  - [ ] usage of project `https://github.com/natefinch/lumberjack`
  - [ ] log level -- integration with zap `https://github.com/uber-go/zap`
* [ ] backlog 
    - [ ] Worker handler error job
    - [ ] Finish async mailer flow
    - [ ] Optimus config 
* [ ] Wanted Feature
    - [ ] go mod cache for images
    - [ ] add grpc protocal
    - [ ] ~~rabbitmq integration~~
    - [ ] docker-compose.yml health check if db ready
    - [ ] my own recovery middleware
        1. email notify developer or sentry
        2. recovery and log with informations like error level 
* [ ] Refactoring
    * [x] refactoring lib file structure  
        Ex: lib/ -> config/ -> config.go
    * [x] refactor lib/asseration name to assertion
    * [ ] Refresh database in sequence
    * [ ] Add setting options for request read and write timeout
    * [ ] Add setting options for run mode
    * [ ] Standardize log format
    * [x] Document flow and new schema
    * [x] Redesign database schema
    * [ ] add migration command
        - [ ] up
        - [ ] down
        - [ ] create
    
---
### Next Jobs
* [ ] Standardize log format, wait for select a library
* [ ] jwt library select 
* [ ] table reschema 
    - add relation between user , email_verify
    - email_verify add column user_id, and foreign key
* [ ] add new helpers function GetDB