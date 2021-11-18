// How our dto (user) will be presented to the user

package users

import (
	"encoding/json"
)

// What will be showed to a public request
type PublicUser struct {
	Id int64 `json:"id"`
	DateCreated string `json:"date_created"`
	Status string `json:"status"`
}

// What will be showed to an internal/private request
type PrivateUser struct {
	Id int64 `json:"id"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Email string `json:"email"`
	DateCreated string `json:"date_created"`
	Status string `json:"status"`
}

func (users Users) Marshall(isPublic bool) []interface{} {
	result := make([]interface{}, len(users))
	for index, user := range users {
		result[index] = user.Marshall(isPublic)
	}
	return result
}

func (user *User) Marshall(isPublic bool) interface{} {
	if isPublic {
		return PublicUser{
			Id: user.Id,
			DateCreated: user.DateCreated,
			Status: user.Status,
		}
	}
	userJson, _ := json.Marshal(user)
	var privateUser PrivateUser
	json.Unmarshal(userJson, &privateUser)
	return privateUser
}