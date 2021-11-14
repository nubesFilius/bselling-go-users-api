package date_utils

import (
	"time"
)

const (
	apiDateLayout = "2021-11-14T11:00:00Z"
)

func GetNow() time.Time {
	return time.Now().UTC()
}

func GetNowString() string {
	
	return GetNow().Format(apiDateLayout)
}