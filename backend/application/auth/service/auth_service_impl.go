package authservice

import (
	"fmt"
	"golang-wails-reactjs/backend/domain/auth"
	"golang-wails-reactjs/backend/domain/user"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type service struct {
	repo auth.Repository
}

func New(r auth.Repository) auth.Service {
	return &service{
		repo: r,
	}
}

func (s *service) Login(login *auth.Login) error {
	result, err := s.repo.FindByEmail(login.Email)
	if err != nil {
		return err
	}

	//get to google sheet users
	resultSheet, err := s.repo.FindSheetByEmail(login.Email)
	if err != nil {
		return err
	}

	if !resultSheet.Status {
		return fmt.Errorf(resultSheet.Message)
	}
	//hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(resultSheet.Data.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	resultSheet.Data.Password = string(hashedPassword)

	// create in table session
	if result == nil {
		err = s.repo.CreateSession(&user.User{
			Name:     resultSheet.Data.Name,
			Email:    resultSheet.Data.Email,
			Password: resultSheet.Data.Password,
			Role:     resultSheet.Data.Role,
			LoginAt:  time.Now().UnixMilli(),
		})
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *service) Logout(email string) error {
	_, err := s.repo.FindByEmail(email)
	if err != nil {
		return err
	}

	// delete in table session
	err = s.repo.DeleteSession(email)
	if err != nil {
		return err
	}
	return nil
}
