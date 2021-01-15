## Schema

* users

|name|type|length|default|index|comment|
|---:|---:|---:|---:|---:|---:|
|id|string|32|N|PK|uuid v4 generate by `{root}/models/users`|
|email|string|128|N|UNIQUE||
|name|string|50|''|||
|created_at|timestamp|||||
|updated_at|timestamp|||||
|deleted_at|timestamp|||||

* email_login

|name|type|length|default|index|comment|
|---:|---:|---:|---:|---:|---:|
|email|string|128|N|PK, foreign:fk_email||
|pwd|string|255|N||hash, nullable|
|verified_at|timestamp|||||
|created_at|timestamp|||||
|updated_at|timestamp|||||

* email_verify

|name|type|length|default|index|comment|
|---:|---:|---:|---:|---:|---:|
|id|uuid||N|PK|uuid|
|email|string|128|N|||
|mail_type|string|64|N||verify, reset|
|verification|smallint||0|N||0:未驗證, 1:已驗證|
|user_id|uuid| | |foreign:fk_user_id||
|created_at|timestamp| | | | |
|updated_at|timestamp| | | | |

* auth_tokens

|name|type|length|default|index|comment|
|---:|---:|---:|---:|---:|---:|
|id|uuid||N|PK|uuid|
|user_id|uuid| | | foreign:fk_user_id| |
|login_way|varchar(64)| | | | |
|revoked|boolean| | | | |
|created_at|timestamp| | | | |
|updated_at|timestamp| | | | |

* images

|name|type|length|default|index|comment|
|---:|---:|---:|---:|---:|---:|
|id|bigint| | |PK|auto_increment|
|access_url|varchar(256)|256 | | | access resource path|
|bucket_name|varchar(32)|32| | | |
|bucket_vendor|varchar(16)|16| | | |
|file_name|varchar(32)|32| | | |
|ext|varchar(8)|8| | | |
|user_id|uuid| | |foreign key| |
|status| smallint|2| | | 0:get_presign_url, 1:uploaded, 2:resized|
