package domain

import "errors"

var (
	ErrTokenDuration = errors.New("invalid token duration format")
	ErrTokenCreation = errors.New("error creating token")
	ErrTokenExpired = errors.New("token expired")
	ErrInvalidToken = errors.New("invalid token")
)