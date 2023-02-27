package dtos

type UpdateUserDTO struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Password  string `json:"password"`
	Bio       string `json:"bio"`
	Photo     string `json:"photo"`
}
