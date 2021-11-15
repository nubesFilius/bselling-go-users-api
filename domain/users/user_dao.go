// Here we'll define how the objects in App interact with DB

package users

import (
	"github.com/nubesFilius/bselling-go-users-api/utils/errors"
	"github.com/nubesFilius/bselling-go-users-api/utils/date_utils"
	"github.com/nubesFilius/bselling-go-users-api/datasources/mysql/users_db"
	"fmt"
	"strings"
)

const(
	indexUniqueEmail = "email_UNIQUE"
	queryInsertUser = "INSERT INTO users(first_name, last_name, email, date_created) VALUES(?, ?, ?, ?);"
)

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
	stmt, err := users_db.Client.Prepare(queryInsertUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	// Have to close the connection when done. Specified in doc of `Prepare()`
	defer stmt.Close()

	user.DateCreated = date_utils.GetNowString()

	insertResult, err := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated)
	if err != nil {
		if strings.Contains(err.Error(), indexUniqueEmail) {
			return errors.NewBadRequestError(fmt.Sprintf("email %s already exists", user.Email))
		}
		return errors.NewInternalServerError(fmt.Sprintf("error when trying to save user: %s", err.Error()))
	}

	// Also possible to do as below. However, it won't provide statement validation prior to executing and it's also less performant.
	// It provides just the result of the executed query.
	// result, err := users_db.Client.Exec(queryInsertUser, user.FirstName, user.LastName, user.Email, user.DateCreated)

	userId, err := insertResult.LastInsertId()
	if err != nil {
		return errors.NewInternalServerError("error when trying to get the last insert id")
	}
	user.Id = userId
	return nil

	// #### LEGACY ####
	// existing_user := usersDB[user.Id]
	// if existing_user != nil {
	// 	if existing_user.Email == user.Email {
	// 		return errors.NewBadRequestError(fmt.Sprintf("email %s already registered", user.Email))
	// 	}
	// 	return errors.NewBadRequestError(fmt.Sprintf("user %d already exists", user.Id))
	// }
	// user.DateCreated = date_utils.GetNowString()
	// usersDB[user.Id] = user
	// return nil
}