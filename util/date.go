package util

import "time"

var ()

func Now() string {
	now := time.Now()
	return now.String()
}
