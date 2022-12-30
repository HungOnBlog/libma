package validators

import "errors"

const (
	passwordInRange8To32 = `^.{8,32}$`
	passwordHasLowercase = `\P{Ll}*$`
	passwordHasUppercase = `\P{Lu}*$`
	passwordHasNumber    = `\P{N}*$`
	passwordHasSymbol    = `\P{S}*$`
)

func IsPasswordValid(password string) error {
	if !IsMatch(password, passwordInRange8To32) {
		return errors.New("password must be between 8 and 32 characters")
	}

	if !IsMatch(password, passwordHasLowercase) {
		return errors.New("password must contain at least one lowercase letter")
	}

	if !IsMatch(password, passwordHasUppercase) {
		return errors.New("password must contain at least one uppercase letter")
	}

	if !IsMatch(password, passwordHasNumber) {
		return errors.New("password must contain at least one number")
	}

	if !IsMatch(password, passwordHasSymbol) {
		return errors.New("password must contain at least one symbol")
	}

	return nil
}
