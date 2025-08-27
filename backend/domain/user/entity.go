package user

type User struct {
	ID       int
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
	LoginAt  int64
}

type UserResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Data    User   `json:"data"`
}
