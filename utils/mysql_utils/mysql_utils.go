package mysql_utils

import (
	"github.com/nubesFilius/bselling-go-users-api/utils/errors"
	"github.com/go-sql-driver/mysql"
	"strings"
	"fmt"
)

const (
	errorNoRows = "no rows in result set"
)

func ParseError(err error) *errors.RestErr {
	// Trying to convert error to MySQLError type
	sqlErr, ok := err.(*mysql.MySQLError)
	if !ok {
		if strings.Contains(err.Error(), errorNoRows) {
			return errors.NewNotFoundError("no record matching given id")
		}
		return errors.NewInternalServerError("error parsing databse response")
	}
	switch sqlErr.Number {
	case 1062:
		return errors.NewBadRequestError(fmt.Sprintf("invalid data"))
	}
	return errors.NewInternalServerError("error processing request")
}