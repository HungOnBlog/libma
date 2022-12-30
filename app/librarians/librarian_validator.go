package librarians

import "github.com/HungOnBlog/libma/core/validators"

func IsLibrarianCreateDtoValid(dto *LibrarianCreateDto) bool {
	if !(dto.Name != "" && dto.Email != "" && dto.Password != "") {
		return false
	}

	if !validators.IsEmailValid(dto.Email) {
		return false
	}

	return true
}

func IsLibrarianUpdateDtoValid(dto *LibrarianUpdateDto) bool {
	if !(dto.Name != "" && dto.Email != "") {
		return false
	}

	if !validators.IsEmailValid(dto.Email) {
		return false
	}

	return true
}

func IsLibrarianLoginDtoValid(dto *LibrarianLoginDto) bool {
	if !(dto.Email != "" && dto.Password != "") {
		return false
	}

	if !validators.IsEmailValid(dto.Email) {
		return false
	}

	return true
}
