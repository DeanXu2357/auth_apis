# Design Blueprint & Todo lists

## Apis

* email register -> done
* email verify -> done
* email activate -> done
* ~~email resend~~
* email password recovery
* email password reset
* user identify api and middleware
* refresh token

## Finished Features
* swagger
* jaeger (integration with opentracing)
* db container health check when app start
* test data factory
* data exist assertion
* event listener 
* migration commands
* smtp integration

## Total todo list
* [ ] backlog
    - [ ] Worker handler error job
    - [ ] Optimus config
    - [ ] modify config to take effect in realtime (viper, watch section) need to think twice about this feature
* [ ] Wanted Feature
    - [ ] go mod cache and docker images integration
    - [ ] add grpc protocal
    - [ ] my own recovery middleware
        1. email notify developer or sentry
        2. recovery and log with informations like error level
* [ ] Refactoring
    * [ ] Refresh database in sequence
    * [ ] Add setting options for request read and write timeout
    * [ ] Add setting options for run mode
* [ ] DevOpts About
    - [ ] design ci/cd flow
    - [ ] write k8s yaml

---
### Next Jobs
* [ ] Standardize log format, wait for select a library
    - [ ] usage of project `https://github.com/natefinch/lumberjack`
    - [ ] log use zap `https://github.com/uber-go/zap`
* [ ] coverage file output
* [ ] test verify api
* [ ] recovery password api
* [ ] reset password api
* [ ] research for line bot login
* [ ] refresh token api swagger
* [ ] check test scopes cover range
* [ ] add postgres ui
* [ ] add create migration command

### Execution sequence
* Normal
    1. add opentracing integration with gorm
    2. finish basic apis and test
       * pic upload and localstack
    4. improve speed of initial container
    5. add grpc support
