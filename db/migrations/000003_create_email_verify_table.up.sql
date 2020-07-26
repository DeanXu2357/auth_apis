CREATE TABLE IF NOT EXISTS email_verify (
    id SERIAL PRIMARY KEY,
    email VARCHAR (128),
    verification SMALLINT,
    created_at TIMESTAMP WITH TIME ZONE,
    updated_at TIMESTAMP WITH TIME ZONE
);

COMMENT ON COLUMN email_verify.verification IS '0:未驗證, 1:已驗證';