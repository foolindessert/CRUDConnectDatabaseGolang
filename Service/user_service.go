package service

import (
	"errors"
	"DATABASECRUD/Entity"
)

type UserIface interface {
	Register(user *entity.User) (*entity.User, error)
}

type UserSvc struct{}

func NewUserSvc() UserIface {
	return &UserSvc{}
}

func (u *UserSvc) Register(user *entity.User) (*entity.User, error) {
	// validasi field field user
	if user.Username == "" {
		return nil, errors.New("username cannot be empty")
	}
	if user.Email == "" {
		return nil, errors.New("email cannot be empty")
	}
	if len(user.Password) < 6 {
		return nil, errors.New("password must be minimum 6 characters")
	}
	if user.Age < 8 {
		return nil, errors.New("age must be greater than 8")
	}
	return user, nil
}
