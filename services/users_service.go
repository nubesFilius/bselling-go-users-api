package services

import (
	"github.com/nubesFilius/bselling-go-users-api/domain/users"
	"github.com/nubesFilius/bselling-go-users-api/utils/errors"
)

func GetUser(userId int64) (*users.User, *errors.RestErr) {
	// If we wanted to validate the in64
	// if userId <= 0 {
	// 	return nil, errors.NewBadRequestError("invalid user id")
	// }
	user := &users.User{Id: userId}
	if err := user.Get(); err != nil {
		return nil, err
	}
	return user, nil
}

func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}
	if err := user.Save(); err != nil {
		return nil, err
	}
	return &user, nil
}