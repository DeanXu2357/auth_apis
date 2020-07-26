CREATE TABLE IF NOT EXISTS users (
    id uuid NOT NULL,
    email VARCHAR (128),
    name VARCHAR (64),
    created_at TIMESTAMP WITH TIME ZONE,
    updated_at TIMESTAMP WITH TIME ZONE,
    PRIMARY KEY (id),
    CONSTRAINT unique_email UNIQUE (email)
)