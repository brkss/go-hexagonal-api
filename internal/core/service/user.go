package service

import "github.com/brkss/dextrace-server/internal/core/port"

/*
	UserService implement port.UserService
	and provide access to cache and user repo service
*/

type UserService struct {
	repo 	port.UserRepository
	cache 	port.CacheRepository	
}