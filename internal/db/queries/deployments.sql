-- name: CreateDeployment :one
INSERT INTO deployments (id, app_id, status, created_at) 
    VALUES (?, ?, ?, now()) 
    RETURNING *;

-- name: UpdateDeployment :exec
UPDATE deployments
    SET status = ?, finished_at = ?
    WHERE id = ?;