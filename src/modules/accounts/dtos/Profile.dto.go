package dtos

type ProfileDTO struct {
	FirstName string  `json:"first_name"`
	LastName  string  `json:"last_name"`
	Email     string  `json:"email"`
	Bio       string  `json:"bio"`
	Photo     *string `json:"photo"`
}
