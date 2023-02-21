package dtos

type CreateUserDTO struct {
	FIRST_NAME string `json:"first_name"`
	LAST_NAME  string `json:"last_name"`
	EMAIL      string `json:"email"`
	PHONE      string `json:"phone"`
	PASSWORD   string `json:"password"`
	BIO        string `json:"bio"`
	PHOTO      string `json:"photo"`
}
