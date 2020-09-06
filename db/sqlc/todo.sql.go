// Code generated by sqlc. DO NOT EDIT.
// source: todo.sql

package db

import (
	"context"
)

const deleteAccount = `-- name: DeleteAccount :exec
DELETE FROM todo
WHERE id = $1
`

func (q *Queries) DeleteAccount(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteAccount, id)
	return err
}

const getAccount = `-- name: GetAccount :one
SELECT id, name, completed, created_at FROM todo
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetAccount(ctx context.Context, id int64) (Todo, error) {
	row := q.db.QueryRowContext(ctx, getAccount, id)
	var i Todo
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Completed,
		&i.CreatedAt,
	)
	return i, err
}

const updateCompleted = `-- name: UpdateCompleted :one
UPDATE todo
SET completed = $2
WHERE id = $1
RETURNING id, name, completed, created_at
`

type UpdateCompletedParams struct {
	ID        int64 `json:"id"`
	Completed bool  `json:"completed"`
}

func (q *Queries) UpdateCompleted(ctx context.Context, arg UpdateCompletedParams) (Todo, error) {
	row := q.db.QueryRowContext(ctx, updateCompleted, arg.ID, arg.Completed)
	var i Todo
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Completed,
		&i.CreatedAt,
	)
	return i, err
}
