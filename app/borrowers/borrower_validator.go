package borrowers

func IsCreateDtoValid(dto *BorrowerDto) bool {
	return dto.Name != "" && dto.Email != "" && dto.Password != ""
}

func IsUpdateDtoValid(dto *BorrowerUpdateDto) bool {
	return dto.Name != "" && dto.Email != "" && dto.Password != ""
}

func IsLoginDtoValid(dto *BorrowerLoginDto) bool {
	return dto.Email != "" && dto.Password != ""
}
