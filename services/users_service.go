package services

import (
	"github.com/nubesFilius/bselling-go-users-api/domain/users"
	"github.com/nubesFilius/bselling-go-users-api/utils/errors"
	"strings"
)

func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	if err := users.Validate(&user); err != nil {
		return nil, err
	}
	return &user, nil
}