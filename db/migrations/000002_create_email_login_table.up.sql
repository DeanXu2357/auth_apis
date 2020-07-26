CREATE TABLE IF NOT EXISTS email_login (
    email VARCHAR (128),
    pwd VARCHAR,
    verifed_at TIMESTAMP WITH TIME ZONE,
    created_at TIMESTAMP WITH TIME ZONE,
    updated_at TIMESTAMP WITH TIME ZONE,
    PRIMARY KEY (email),
    CONSTRAINT fk_email
        FOREIGN KEY (email)
            REFERENCES users (email) ON DELETE CASCADE ON UPDATE CASCADE
)