package port

import "github.com/brkss/dextrace-server/internal/core/domain"

// TokenService is an interface for interacting with token-related buisness logic
type TokenService interface {
	// CreateToken create token new token for a given user 
	CreateToken(user *domain.User) (string, error)
	// verifyToken verify a given  token and return it's payload
	VerifyToken(token string)(*domain.TokenPayload, error)
}

