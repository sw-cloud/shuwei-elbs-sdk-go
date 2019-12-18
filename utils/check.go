package utils

import "strings"

func IsEmptyString(str string) bool  {
	if strings.TrimSpace(str) == "" {
		return true
	}
	return false
}
