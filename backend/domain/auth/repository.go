package auth

import "golang-wails-reactjs/backend/domain/user"

type Repository interface {
	FindByEmail(email string) (*user.User, error)
	FindSheetByEmail(email string) (*user.UserResponse, error)
	CreateSession(*user.User) error
}
