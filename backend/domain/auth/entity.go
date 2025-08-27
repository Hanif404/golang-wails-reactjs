package auth

type Login struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

type Logout struct {
	Email string `json:"email" validate:"required,email"`
}
