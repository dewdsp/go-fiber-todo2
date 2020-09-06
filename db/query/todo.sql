-- name: GetAccount :one
SELECT * FROM todo
WHERE id = $1 LIMIT 1;

-- name: DeleteAccount :exec
DELETE FROM todo
WHERE id = $1;

-- name: UpdateCompleted :one
UPDATE todo
SET completed = $2
WHERE id = $1
RETURNING *;