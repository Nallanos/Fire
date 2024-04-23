-- name: CreateDeployment :one
INSERT INTO deployments (id, app_id, status, created_at) 
    VALUES (?, ?, ?, CURRENT_TIMESTAMP) 
    RETURNING *;

-- name: UpdateDeployment :exec
UPDATE deployments
    SET status = ?, finished_at = ?
    WHERE id = ?;

-- name: GetDeployment :one
SELECT * FROM deployments WHERE app_id = ? LIMIT 1;