package users

import (
	"github.com/nubesFilius/bselling-go-users-api/utils"
	"strings"
)

type User struct {
	Id int64
	Firstname string
	Lastname string
	Email string
	DateCreated string
}

func (user *User) Validate() *errors.RestErr {
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email == "" {
		return errors.NewBadRequestError("invalid email address")
	}
	return nil
}