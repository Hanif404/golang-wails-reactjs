package auth

type Service interface {
	Login(login *Login) error
}
