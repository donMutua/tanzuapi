// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0

package db

import (
	"database/sql"
	"time"
)

type Event struct {
	ID        int64        `json:"id"`
	EventName string       `json:"event_name"`
	CreatedAt sql.NullTime `json:"created_at"`
	About     string       `json:"about"`
	EventDate time.Time    `json:"event_date"`
	Mode      string       `json:"mode"`
	Cost      string       `json:"cost"`
	StartTime time.Time    `json:"start_time"`
	EndTime   time.Time    `json:"end_time"`
}

type PaymentHistory struct {
	ID                int32          `json:"id"`
	TicketID          sql.NullInt64  `json:"ticket_id"`
	Status            sql.NullString `json:"status"`
	TransactionID     sql.NullString `json:"transaction_id"`
	TransactionType   sql.NullString `json:"transaction_type"`
	Recipient         sql.NullString `json:"recipient"`
	ParentTransaction sql.NullString `json:"parent_transaction"`
	Email             sql.NullString `json:"email"`
	Recieved          sql.NullTime   `json:"recieved"`
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
	ID            int32         `json:"id"`
	EventID       sql.NullInt64 `json:"event_id"`
	Attendee      string        `json:"attendee"`
	PaymentStatus string        `json:"payment_status"`
	CreatedAt     sql.NullTime  `json:"created_at"`
}

type User struct {
	Username          string    `json:"username"`
	HashedPassword    string    `json:"hashed_password"`
	FirstName         string    `json:"first_name"`
	LastName          string    `json:"last_name"`
	Email             string    `json:"email"`
	Role              string    `json:"role"`
	PasswordChangedAt time.Time `json:"password_changed_at"`
	CreatedAt         time.Time `json:"created_at"`
}
