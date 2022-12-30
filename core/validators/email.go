package validators

import "regexp"

// Validate email address
//
// Example:
//
//	IsEmailValid(" ") // false
//	IsEmailValid("test") // false
//	IsEmailValid("test@") // false
//	IsEmailValid("test@test") // false
//	IsEmailValid("test@test.") // false
//	IsEmailValid("test@test.test") // true
func IsEmailValid(email string) bool {
	emailValidationPattern := `^([a-zA-Z0-9_\.\-])+\@(([a-zA-Z0-9\-])+\.)+([a-zA-Z0-9]{2,4})+$`

	return IsMatch(email, emailValidationPattern)
}

func IsMatch(str string, pattern string) bool {
	regex := regexp.MustCompile(pattern)

	return regex.MatchString(str)
}
