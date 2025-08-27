package auth

type Service interface {
	Login(login *Login) error
	Logout(email string) error
}
