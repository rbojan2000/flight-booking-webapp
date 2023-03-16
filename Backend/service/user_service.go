package service

import (
	"flightbooking-app/model"
	"flightbooking-app/repo"
)

type UserService struct {
	UserRepo *repo.UserRepository
}

func (service *UserService) Create(user *model.User) error {
	err := service.UserRepo.Create(user)
	if err != nil {
		return err
	}
	return nil
}
