package borrowers

func IsCreateDtoValid(dto *BorrowerDto) bool {
	return dto.Name != "" && dto.Email != "" && dto.Password != ""
}
