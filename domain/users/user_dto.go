// Here we define the objects that will be
// transfered between app and DB.

package users

type User struct {
	Id int64 `json:"id"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Email string `json:"email"`
	DateCreated string `json:"date_created"`
}

// ### No need to Validate, as it is handled by the MySQL DB.
// func (user *User) Validate() *errors.RestErr {
// 	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
// 	if user.Email == "" {
// 		return errors.NewBadRequestError("invalid email address")
// 	}
// 	return nil
// }