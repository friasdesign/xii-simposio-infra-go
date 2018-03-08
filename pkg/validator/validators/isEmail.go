package validators

import (
	"regexp"
)

// IsEmail validates whether the string i is a valid Email address.
func IsEmail(i string) bool {
	r, err := regexp.MatchString("^[a-zA-Z0-9_.+-]+@[a-zA-Z0-9-]+\\.[a-zA-Z0-9-.]+$", i)
	if err != nil {
		panic(err)
	}
	return r
}