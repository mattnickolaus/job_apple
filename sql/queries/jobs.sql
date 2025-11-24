-- name: CreateJob :one
INSERT INTO jobs (title, company, applied_at, updated_at, link, status)
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6
)
RETURNING *;
