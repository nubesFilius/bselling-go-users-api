// Here we'll define how the objects in App interact with DB

package users

import (
	"github.com/nubesFilius/bselling-go-users-api/utils/date_utils"
	"github.com/nubesFilius/bselling-go-users-api/utils/errors"
	"fmt"
)

// Mock DB for testing
var(usersDB = make(map[int64]*User))

func (user *User) Get() *errors.RestErr {
	result := usersDB[user.Id]
	if result == nil {
		return errors.NewNotFoundError(fmt.Sprint("user %d not found", user.Id))
	}
	user.Id = result.Id
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.Email = result.Email
	user.DateCreated = result.DateCreated
	return nil
}

func (user *User) Save() *errors.RestErr {
	existing_user := usersDB[user.Id]
	// If existing_user_user not empty (it exists)
	if existing_user != nil {
		if existing_user.Email == user.Email {
			return errors.NewBadRequestError(fmt.Sprintf("email %s already registered", user.Email))
		}
		return errors.NewBadRequestError(fmt.Sprintf("user %d already exists", user.Id))
	}
	user.DateCreated = date_utils.GetNowString()
	usersDB[user.Id] = user
	return nil
}