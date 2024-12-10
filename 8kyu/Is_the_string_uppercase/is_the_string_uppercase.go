package Is_the_string_uppercase

import (
	"unicode"
)

type MyString string

func (s MyString) IsUpperCase() bool {
	for _, char := range s {
		if unicode.IsLetter(char) && !unicode.IsUpper(char) {
			return false
		}
	}
	return true
}
