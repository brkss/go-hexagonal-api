package domain

import (
	"github.com/google/uuid"
)

// TokenPayload is an entity that represent the payload of a token
type TokenPayload struct {
	ID uuid.UUID
	UserID string 
}