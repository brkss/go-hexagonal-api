package domain

import "time"

// User is an entity that represent user
type User struct {
	ID string
	Name string
	Email string
	Password string
	CreatedAt time.Time
	UpdatedAt time.Time
}