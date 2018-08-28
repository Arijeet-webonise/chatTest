
-- +migrate Up

CREATE TABLE access_control (
    id SERIAL PRIMARY KEY,
    route_name VARCHAR(255) UNIQUE NOT null,
    route_path VARCHAR(255) UNIQUE NOT null
);

CREATE TABLE access_control_mapping (
    id SERIAL PRIMARY KEY,
    access_id int REFERENCES access_control(id) NOT null,
    role roleType NOT NULL,
    permission bool NOT null,
    CONSTRAINT accessid_role UNIQUE (access_id, role)
);


-- +migrate Down

DROP TABLE access_control;
DROP TABLE access_control_mapping;
