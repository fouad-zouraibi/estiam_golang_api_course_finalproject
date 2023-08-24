package services

import (
	"context"

	"github.com/phramos07/finalproject/repos"
	"github.com/phramos07/finalproject/types"
)

type UserService interface {
	CreateNewUser(context.Context, *types.User) error
	GetUser(context.Context, string) (*types.User, error)
}

type userServiceImpl struct {
	repo repos.UserRepository
}

func NewUserService(repo repos.UserRepository) UserService {
	return &userServiceImpl{repo: repo}
}

/*
TODO: Add service that ENCRYPTS PASSWORD and STORES IN DB

Tips
- Use bcrypt package!
- Save bcrypt secret on .env and load it in App configuration!
- Inject app configuration (bcrypt secret) into here (user service)
*/
func (u *userServiceImpl) CreateNewUser(c context.Context, user *types.User) error {
	//TODO: Logic to create user - hash the password !!!

	return nil
}

/*REMOVE THIS SERVICE*/
func (u *userServiceImpl) GetUser(c context.Context, id string) (*types.User, error) {
	return u.repo.GetUser(c, id)
}
