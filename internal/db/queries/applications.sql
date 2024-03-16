-- name: ListApplications :many
SELECT * FROM applications;

-- name: GetApplication :one
SELECT * FROM applications WHERE id = ? LIMIT 1;


-- name: CreateApplication :one
INSERT INTO applications (id, name, image, port)
VALUES (?, ?, ?, ?)
RETURNING *;

-- name: UpdateApplication :one
UPDATE applications
SET name = ?, image = ?, port = ?
WHERE id = ?
RETURNING *;

-- name: DeleteApplication :exec
DELETE FROM applications WHERE id = ?;


