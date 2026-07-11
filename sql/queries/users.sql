-- name: CreateUser :exec
INSERT INTO users (name, email, hashed_password, created)
VALUES($1, $2, $3, NOW());