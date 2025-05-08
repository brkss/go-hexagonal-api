package service

import (
	"context"

	"github.com/brkss/dextrace-server/internal/core/domain"
	"github.com/brkss/dextrace-server/internal/core/port"
	"github.com/brkss/dextrace-server/internal/core/utils"
)

/*
	UserService implement port.UserService
	and provide access to cache and user repo service
*/

type UserService struct {
	repo 	port.UserRepository
	cache 	port.CacheRepository	
}

func NewUserService(repo port.UserRepository, cache port.CacheRepository) *UserService {
	return &UserService{
		repo,
		cache,
	}
}

func (uc *UserService) Register(ctx context.Context, user *domain.User) (*domain.User, error) {

	hashed, err := utils.HashPassword(user.Password) 
	if err != nil {
		return nil, domain.ErrInternal
	}

	return user;
}