package utils

import "strings"

func LowerCamelCase(v string) string {
	s := strings.Split(v, "")
	s[0] = strings.ToLower(s[0])

	res := strings.Join(s, "")

	return res
}
