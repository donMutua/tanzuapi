// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0

package db

import (
	"database/sql"
)

type Event struct {
	ID        int64          `json:"id"`
	UserID    sql.NullInt64  `json:"user_id"`
	EventName string         `json:"event_name"`
	CreatedAt sql.NullTime   `json:"created_at"`
	About     sql.NullString `json:"about"`
}

type Speaker struct {
	ID        int32         `json:"id"`
	EventID   sql.NullInt64 `json:"event_id"`
	FirstName string        `json:"first_name"`
	LastName  string        `json:"last_name"`
	CreatedAt sql.NullTime  `json:"created_at"`
	Bio       string        `json:"bio"`
}

type Ticket struct {
	ID        int32          `json:"id"`
	EventID   sql.NullInt64  `json:"event_id"`
	Price     sql.NullInt64  `json:"price"`
	StartDate sql.NullTime   `json:"start_date"`
	EndDate   sql.NullTime   `json:"end_date"`
	Mode      sql.NullString `json:"mode"`
	CreatedAt sql.NullTime   `json:"created_at"`
}

type User struct {
	ID        int64        `json:"id"`
	FirstName string       `json:"first_name"`
	LastName  string       `json:"last_name"`
	Email     string       `json:"email"`
	Role      string       `json:"role"`
	CreatedAt sql.NullTime `json:"created_at"`
}