package user

type User struct {
	ID       int
	Name     string
	Email    string
	Password string
	Role     string
	LoginAt  int64
}

type UserResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Data    User   `json:"data"`
}
