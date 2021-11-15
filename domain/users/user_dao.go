// Here we'll define how the objects in App interact with DB

package users

import (
	"github.com/nubesFilius/bselling-go-users-api/utils/date_utils"
	"github.com/nubesFilius/bselling-go-users-api/utils/errors"
	"github.com/nubesFilius/bselling-go-users-api/datasources/mysql/users_db"
	"fmt"
)

// Mock DB for testing
var(usersDB = make(map[int64]*User))

func (user *User) Get() *errors.RestErr {
	if err := users_db.Client.Ping(); err != nil {
		panic(err)
	}

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
	fmt.Println(user)
	user.DateCreated = date_utils.GetNowString()
	fmt.Println(user)
	usersDB[user.Id] = user
	return nil
}