package service

import (
	"github.com/ivo100/go-practice/userservice/internal/core"
	"github.com/ivo100/go-practice/userservice/internal/repository"
	svc "github.com/ivo100/go-practice/userservice/pkg"
)

// UserServiceImpl provides a REST API UserService implementation for working with users.
type UserServiceImpl struct {
	repo core.UserRepository
}

func NewUserService() svc.UserService {
	return &UserServiceImpl{repo: repository.NewMemoryUserRepository()}
}
func (s UserServiceImpl) AddUser(user svc.User) (*svc.User, error) {
	// convert service user to dto
	dtoUser, err := convertToDtoUser(&user)
	if err != nil {
		return nil, err
	}
	// add user to repository
	dto, err := s.repo.SaveUser(*dtoUser)
	if err != nil {
		return nil, err
	}
	// convert dto to service user
	return convertToServiceUser(dto)
}

func (s UserServiceImpl) GetUser(id string) (*svc.User, error) {
	user, err := s.repo.GetUser(id)
	if err != nil {
		return nil, err
	}
	return convertToServiceUser(user)
}

func convertToDtoUser(user *svc.User) (*core.UserDto, error) {
	if user == nil {
		return nil, nil
	}
	res := &core.UserDto{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
	}
	return res, nil
}

func convertToServiceUser(user *core.UserDto) (*svc.User, error) {
	if user == nil {
		return nil, nil
	}
	res := &svc.User{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
	}
	return res, nil
}
