// Code generated by sqlc. DO NOT EDIT.

package db

import (
	"time"
)

type Todo struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Completed bool      `json:"completed"`
	CreatedAt time.Time `json:"created_at"`
}
