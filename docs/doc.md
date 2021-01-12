# Design Blueprint

## Apis

* email register -> done
* email verify -> done 
* email activate -> done
* ~~email resend~~
* email password recovery
* email password reset
* user identify api and middleware
* refresh token

## Total todo list
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
* [x] verify api
* [ ] Log file writing
  - [ ] usage of project `https://github.com/natefinch/lumberjack`
  - [ ] log level -- integration with zap `https://github.com/uber-go/zap`
  - [ ] Standardize log format by zap
* [ ] backlog 
    - [ ] Worker handler error job
    - [x] Finish async mailer flow
    - [ ] Optimus config 
* [ ] Wanted Feature
    - [ ] go mod cache and docker images integration
    - [ ] add grpc protocal
    - [x] docker-compose.yml health check if db ready
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
    * [x] Document flow and new schema
    * [x] Redesign database schema
    * [ ] add migration command
        - [ ] up
        - [ ] down
        - [ ] create
* [ ] DevOpts About
    - [ ] design ci/cd flow
    - [ ] write k8s yaml
    
---
### Next Jobs
* [ ] Standardize log format, wait for select a library
* [ ] coverage file output
* [ ] test verify api
* [ ] verify token middleware
* [ ] refresh token api 
* [ ] recovery password api
* [ ] reset password api 
* [ ] research for line bot login

### Execution sequence
* Normal 
    1. finish basic apis and test
        * pic upload and localstack
    2. add grpc support
    3. swagger
    4. open tracing support
    5. add migration command and fix README.md about initial project
    
* DevOpts
    1. design ci/cd flow
