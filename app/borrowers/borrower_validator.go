package borrowers

import "github.com/HungOnBlog/libma/core/validators"

func IsCreateDtoValid(dto *BorrowerDto) bool {
	if !(dto.Name != "" && dto.Email != "" && dto.Password != "") {
		return false
	}

	if !validators.IsEmailValid(dto.Email) {
		return false
	}

	return true
}

func IsUpdateDtoValid(dto *BorrowerUpdateDto) bool {
	if !(dto.Name != "" && dto.Email != "") {
		return false
	}

	if !validators.IsEmailValid(dto.Email) {
		return false
	}

	return true
}

func IsLoginDtoValid(dto *BorrowerLoginDto) bool {
	if !(dto.Email != "" && dto.Password != "") {
		return false
	}

	if !validators.IsEmailValid(dto.Email) {
		return false
	}

	return true
}
