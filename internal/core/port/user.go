package port

import (
	"context"

	"github.com/brkss/dextrace-server/internal/core/domain"
)




type UserRepository interface {
	// CreateUser creates a new user in the database 
	CreateUser(ctx context.Context, user *domain.User) (*domain.User, error);
	// GetUserById gets user by id 
	GetUserById(ctx context.Context, id int64)(*domain.User, error)
	// GetUserByEmail gets user by it's email 
	GetUserByEmail(ctx context.Context, email string) (*domain.User, error)
}


type UserService interface {
	// Register registers a new user 
	Register(ctx context.Context, user *domain.User) (*domain.User, error)
	// GetUser get user by id 
	GetUser(ctx context.Context, id int64) (*domain.User, error)
	// GetUserByEmail get user by email address 
}