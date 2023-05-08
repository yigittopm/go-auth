package service

import (
	"go-auth/pkg/model"
	"go-auth/pkg/repository/user"
)

type UserService struct {
	UserRepository *user.Repository
}

func NewUserService(u *user.Repository) UserService {
	return UserService{
		UserRepository: u,
	}
}

func (u *UserService) All() ([]model.User, error) {
	return u.UserRepository.All()
}

func (u *UserService) FindById(id uint) (*model.User, error) {
	return u.UserRepository.FindById(id)
}

func (u *UserService) Save(user *model.User) (*model.User, error) {
	return u.UserRepository.Save(user)
}

func (u *UserService) Delete(id uint) error {
	return u.UserRepository.Delete(id)
}

func (u *UserService) Migrate() error {
	return u.UserRepository.Migrate()
}
