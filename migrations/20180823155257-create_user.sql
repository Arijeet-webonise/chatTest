
-- +migrate Up

CREATE TYPE roleType AS ENUM ('SUPERADMIN', 'ADMIN', 'USER');

CREATE TABLE portal_user (
    id SERIAL PRIMARY KEY,
    first_name VARCHAR(255) NOT NULL,
    last_name VARCHAR(255) NOT NULL,
    email_id VARCHAR(255) NOT NULL CONSTRAINT unq_ename UNIQUE,
    password VARCHAR(255) NOT NULL,
    role roleType NOT NULL DEFAULT 'USER',
    modified_at TIMESTAMP NOT NULL DEFAULT now(),
    created_at TIMESTAMP NOT NULL DEFAULT now()
);


-- +migrate Down
DROP TABLE portal_user;
DROP TYPE roleType;
