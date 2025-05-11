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

func (us *UserService) Register(ctx context.Context, user *domain.User) (*domain.User, error) {

	hashed, err := utils.HashPassword(user.Password) 
	if err != nil {
		return nil, domain.ErrInternal
	}

	user.Password = hashed
	user, err = us.repo.CreateUser(ctx, user)
	if err != nil {
		if err == domain.ErrConflictingData {
			return nil, err
		}
		return nil, domain.ErrInternal
	}

	cacheKey := utils.GenerateCacheKey("user", user.ID)
	userSerialized, err := utils.Serialize(user)
	if err != nil {
		return nil, domain.ErrInternal;
	}

	err = us.cache.Set(ctx, cacheKey, userSerialized, 0)
	if err != nil {
		return nil, domain.ErrInternal
	}

	// caching ...
	return user, nil;
}

func (us *UserService) GetUser(ctx context.Context, id int64) (*domain.User, error) {
	var user *domain.User;

	cachedKey := utils.GenerateCacheKey("user", id)
	cachedUser, err := us.cache.Get(ctx, cachedKey)
	if err != nil {
		err := utils.Deserialize(cachedUser, &user)
		if err != nil {
			return nil, domain.ErrInternal
		}
		return user, nil
	}

	user, err = us.repo.GetUserById(ctx, id)
	if err != nil {
		if err == domain.ErrNoDataFound {
			return nil, err
		}
		return nil, domain.ErrInternal
	}

	userSerialized, err := utils.Serialize(user)
	if err != nil {
		return nil, domain.ErrInternal
	}

	err = us.cache.Set(ctx, cachedKey, userSerialized, 0)
	if err != nil {
		return nil, domain.ErrInternal
	}

	return user, nil
}