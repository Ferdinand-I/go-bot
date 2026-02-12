package misc

import "strings"

func Must(err error) {
	if err != nil {
		panic(err)
	}
}

func GetCallbackQueryPrefix(s string) string {
	data := strings.Split(s, ":")
	return data[0]
}
