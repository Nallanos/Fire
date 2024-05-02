-- name: CreateDeployment :one
INSERT INTO deployments (id, app_id, status, created_at) 
    VALUES (?, ?, ?, CURRENT_TIMESTAMP) 
    RETURNING *;

-- name: UpdateDeployment :exec
UPDATE deployments
    SET status = ?, finished_at = ?
    WHERE app_id = ?;

-- name: GetLatestDeployment :one
SELECT * FROM deployments
WHERE app_id = ?
ORDER BY created_at DESC
LIMIT 1;

-- name: ListApplicationDeployments :many
SELECT * FROM deployments WHERE app_id = ?;