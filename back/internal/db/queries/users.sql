-- name: GetUserByID :one
SELECT * FROM users WHERE id = ?;

-- name: GetUserByEmail :one
SELECT * FROM users WHERE email = ?;

-- name: CreateUser :exec
INSERT INTO users (id, email, password, name) VALUES (?, ?, ?, ?);

-- name: UpdateUserName :exec
UPDATE users SET name = ? WHERE id = ?;

-- name: DeleteUser :exec
DELETE FROM users WHERE id = ?;

-- name: CreateSession :exec
INSERT INTO tokens (user_id, token, expires_at) VALUES (?, ?, ?);

-- name: GetSession :one
SELECT * FROM tokens WHERE token = ?;

-- name: GetSessionByUserId :one
SELECT * FROM tokens WHERE user_id = ?;

-- name: DeleteSession :exec
DELETE FROM tokens WHERE token = ?;