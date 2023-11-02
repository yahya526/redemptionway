package util

func TrimEndSingleQuotationMark(str string) string {
	if len(str) == 0 {
		return str
	}
	if str[len(str)-1] == '\'' {
		str = str[0 : len(str)-1]
	}
	return str
}

func TrimStartSwapLine(str string) string {
	if len(str) < 2 {
		return str
	}
	if str[0:2] == "\n" {
		return str[2:]
	}
	if len(str) < 4 {
		return str
	}
	if str[0:4] == "\r\n" {
		return str[4:]
	}
	return str
}
