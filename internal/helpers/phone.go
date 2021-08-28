package helpers

import "regexp"

func IsValidPhone(phone string) bool {

	isValid, _ := regexp.MatchString("(05|5)[0-9][0-9][1-9]([0-9]){6}", phone)

	return isValid
}
