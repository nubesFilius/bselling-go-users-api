package date_utils

import (
	"time"
)

const (
	apiDateLayout = "Jan 2 15:04:05 2006 MST"
)

func GetNow() time.Time {
	return time.Now().UTC()
}

func GetNowString() string {
	return GetNow().Format(apiDateLayout)
}