# Register flows Spec

## Apis
* Register
    * method: POST
    * path: `{domain}/api/v1/register`
    * json parameters: `{"name": "dean", "email": "dean@example.com", "password": "password"}`
    * json response: 
        * success: http code (200), `{"status_code":"200", "msg":"success"}`
        * user not found: http code (400), `{"status_code":"4004", "msg":"user not found"}`
        * user already exists: http code (400), `{"status_code":"4009", "msg":"user already exists"}`
        * user already exists, but not certificate yet: http code (400), `{"status_code":"40092", "msg":"user already exists, but not certificate yet"}`
    
* Resend Registration (deprecated)
    * memo : need to implement resend limitation 
    * method: POST
    * path: `{domain}/api/v1/register/resend`
    * json parameters: `{"email": "dean@example.com"}`
    * json response: 
        * success: http code (200), `{"status_code":"200", "msg":"success"}`
        * registration not found: http code (400), `{"status_code":"4004", "msg":"registration not exists"}`
        * resend too often: http code (429), `{"status_code":"429", "msg":"too many requests"}`
        
* Activate Registration
    * method: POST
    * path: `{domain}/api/v1/register/activate`
    * json parameters: `{"activation_code": "iiiiiiiiiiiiii"}`
    * json response:
        * success: http code (200), `{"status_code":"200", "msg":"success"}`
        * unauthorized activation_code
        * activation_code expires
         
## Flow

* Registration
    1. check if email already registered
    2. produce activation code
    3. insert data into table email_certificates 
    4. send registration email
    
* Register by email
    1. request api `/api/v1/email/register`
    2. model user create
    3. model email_login create
    4. model email_verify create
    5. send email
        - url = `/api/v1/email/activate?token=` + token(use email_verify id as id to generate jwt token)
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
        - url = `/api/v1/email/activate?token=` + token(use email_verify id as id to generate jwt token)    
    4. request reset password api `/api/v1/email/reset`
    5. query email_verify by id
    6. verify
        1. email
        2. type
        3. created_at
    7. update email_login
    8. delete email_verify
    9. redirect to login page (setting in config file)

    
## Remarks
* To reuse table email_certificates in reset password flow . 
* Remember to write a cronjob delete expired email_certificates data raw everyday .
