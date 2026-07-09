-- +goose up
CREATE TABLE IF NOT EXISTS users(
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    hashed_password CHAR(60) NOT NULL,
    created TIMESTAMPTZ NOT NULL
);
ALTER TABLE users
ADD CONSTRAINT users_uc_email UNIQUE (email);
-- +goose down
DROP TABLE users;