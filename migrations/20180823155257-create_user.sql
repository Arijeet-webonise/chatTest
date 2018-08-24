
-- +migrate Up

CREATE TABLE portal_user (
    id SERIAL PRIMARY KEY,
    first_name VARCHAR(255),
    last_name VARCHAR(255),
    email_id VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    modified_at TIMESTAMP NOT NULL DEFAULT now(),
    created_at TIMESTAMP NOT NULL DEFAULT now()
);

-- +migrate Down
DROP TABLE portal_user;
