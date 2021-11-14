package services

import (
	"github.com/nubesFilius/bselling-go-users-api/domain/users"
	"github.com/nubesFilius/bselling-go-users-api/utils"
)

func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}
	if err := user.Save(); err != nil {
		return nil, err
	}
	return &user, nil
}