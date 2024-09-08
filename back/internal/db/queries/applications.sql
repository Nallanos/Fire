-- name: ListApplications :many
SELECT * FROM applications WHERE user_id = ?;

-- name: GetApplication :one
SELECT * FROM applications WHERE id = ? LIMIT 1;

-- name: CreateApplication :one
INSERT INTO applications (id, name, image, port, status, user_id)
VALUES (?, ?, ?, ?, ?, ?)
RETURNING *;

-- name: UpdateApplication :exec
UPDATE applications
SET name = ?, image = ?, port = ?, status = ?
WHERE id = ?;

-- name: DeleteApplication :exec
DELETE FROM applications WHERE id = ?;


