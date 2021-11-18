package date_utils

import (
	"time"
)

const (
	apiDateLayout = "Jan 2 15:04:05 2006 MST"
	apiDBLayout   = "2021-11-16 15:00:00"
)

func GetNow() time.Time {
	return time.Now().UTC()
}

func GetNowString() string {
	return GetNow().Format(apiDateLayout)
}

func GetNowDBFormat() string {
	return GetNow().Format(apiDBLayout)
}
