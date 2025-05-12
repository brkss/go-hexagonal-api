package domain

import "errors"

var (
	ErrTokenDuration = errors.New("invalid token duration format")
	ErrTokenCreation = errors.New("error creating token")
	ErrTokenExpired = errors.New("token expired")
	ErrInvalidToken = errors.New("invalid token")
	ErrConflictingData = errors.New("data conflicts with existing data in unique column")
	ErrNoDataFound = errors.New("no data was found")
	ErrInternal = errors.New("internal server error")
	ErrInvalidCredentials = errors.New("invalid user credentials")
	ErrEmptyAuthorizationHeader = errors.New("empty authorization key")
	ErrInvalidAuthorizationHeader = errors.New("invalid authorization header")
	ErrInvalidAuthorizationType = errors.New("invalid authorization type")
)