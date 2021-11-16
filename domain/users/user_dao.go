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
	errorNoRows = "no rows in result set"
	queryInsertUser = "INSERT INTO users(first_name, last_name, email, date_created) VALUES(?, ?, ?, ?);"
	queryGetUser = "SELECT id, first_name, last_name, email, date_created FROM users WHERE id=?;"
)

func (user *User) Get() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryGetUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	result := stmt.QueryRow(user.Id)
	if err := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated); err != nil {
		if strings.Contains(err.Error(), errorNoRows) {
			return errors.NewNotFoundError(fmt.Sprintf("user %d not found", user.Id))
		}
		return errors.NewInternalServerError(fmt.Sprintf("error when trying to get user %d: %s", user.Id, err.Error()))
	}

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