package service

import (
	"context"

	"github.com/brkss/dextrace-server/internal/core/domain"
	"github.com/brkss/dextrace-server/internal/core/port"
	"github.com/brkss/dextrace-server/internal/core/utils"
)

type AuthService struct {
	repo port.UserRepository
	ts port.TokenService
}


func NewAuthService(repo port.UserRepository, token port.TokenService) *AuthService {
	return &AuthService{
		repo,
		token,
	}
}


func (as *AuthService)Login(ctx context.Context, email, password string) (string, error) {
	user, err := as.repo.GetUserByEmail(ctx, email)

	if err != nil {
		if err == domain.ErrNoDataFound {
			return "", domain.ErrInvalidCredentials	
		}
	}


	err = utils.ValidatePassword(password, user.Password)
	if err != nil {
		return "", domain.ErrInvalidCredentials
	}

	accessToken, err := as.ts.CreateToken(user)
	if err != nil {
		return "", domain.ErrTokenCreation
	}

	return accessToken, nil
}