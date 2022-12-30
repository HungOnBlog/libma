package borrowers

type BorrowerDto struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type BorrowerGetDto struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type BorrowerUpdateDto struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type BorrowerLoginDto struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
