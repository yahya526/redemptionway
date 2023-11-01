package util

import "time"

var (
	DateTimeSimple = "20060102150405"
)

func NowStr() string {
	return time.Now().Format(time.DateTime)
}

func NowStrSimple() string {
	return time.Now().Format(DateTimeSimple)
}

func Unix() {
}
