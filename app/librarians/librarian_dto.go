package librarians

type LibrarianCreateDto struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LibrarianUpdateDto struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LibrarianDto struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type LibrarianLoginDto struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
