CREATE TABLE IF NOT EXISTS auth_tokens (
    id uuid NOT NULL,
    user_id uuid,
    login_way VARCHAR (64),
    revoked BOOLEAN NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE,
    updated_at TIMESTAMP WITH TIME ZONE,
    PRIMARY KEY (id),
    CONSTRAINT fk_user_id
        FOREIGN KEY (user_id)
            REFERENCES users (id) ON DELETE CASCADE
)