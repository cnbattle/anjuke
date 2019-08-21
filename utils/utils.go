package utils

import "time"

func GetYmd() string {
	return time.Now().Format("2006") +
		time.Now().Format("01") +
		time.Now().Format("02")
}
