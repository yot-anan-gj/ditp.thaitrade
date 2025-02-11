package stringutil

import "strings"

//IsEmptyString return true if empty string
func IsEmptyString(input string) bool {
	if strings.TrimSpace(input) == "" || len(input) == 0 {
		return true
	}
	return false
}

//IsNotEmptyString return true if notempty string
func IsNotEmptyString(input string) bool {
	if strings.TrimSpace(input) != "" && len(input) > 0 {
		return true
	}
	return false
}
