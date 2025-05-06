package domain

import "time"

// User is an entity that represent user
type User struct {
	ID uint64
	Name string
	Email string
	CreatedAt time.Time
	UpdatedAt time.Time
}