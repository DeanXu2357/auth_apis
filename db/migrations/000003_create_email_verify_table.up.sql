CREATE TABLE IF NOT EXISTS email_verify (
    id uuid NOT NULL,
    email VARCHAR (128),
    mail_type VARCHAR (64),
    verification SMALLINT,
    user_id uuid,
    created_at TIMESTAMP WITH TIME ZONE,
    updated_at TIMESTAMP WITH TIME ZONE,
    PRIMARY KEY (id)
);

COMMENT ON COLUMN email_verify.verification IS '0:未驗證, 1:已驗證';
COMMENT ON COLUMN email_verify.mail_type IS 'verify:驗證信箱, reset:重置密碼';