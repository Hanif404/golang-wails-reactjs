package authrepository

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"golang-wails-reactjs/backend/domain/auth"
	"golang-wails-reactjs/backend/domain/user"
	"golang-wails-reactjs/backend/pkg/google"
)

type repo struct {
	db *sql.DB
}

func New(db *sql.DB) auth.Repository {
	return &repo{
		db,
	}
}

func (r *repo) FindByEmail(email string) (*user.User, error) {
	var user user.User

	// Execute the query
	result := r.db.QueryRow("SELECT id, name, loginAt FROM session where email = ?", email)
	err := result.Scan(&user.ID, &user.Name, &user.LoginAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // not found
		}
		return nil, err
	}
	return &user, nil
}

func (r *repo) FindSheetByEmail(email string) (*user.UserResponse, error) {
	var user user.UserResponse

	result, err := google.GetResponse("AKfycbzmoNWHUHEEovXld1gzt7fcGfQPCmaoNbIfTL9Nl8KJr6e6O6aaNNWfd28EAh7H8u-7kw", "?email="+email)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(result, &user)
	if err != nil {
		fmt.Printf("Error unmarshaling JSON: %v\n", err)
		return nil, err
	}

	return &user, nil
}

func (r *repo) CreateSession(user *user.User) error {
	_, err := r.db.Exec("INSERT INTO session (email, name, password, role, loginAt) VALUES (?, ?, ?, ?, ?)", user.Email, user.Name, user.Password, user.Role, user.LoginAt)
	if err != nil {
		return err
	}
	return nil
}

func (r *repo) DeleteSession(email string) error {
	_, err := r.db.Exec("DELETE FROM session WHERE email = ?", email)
	if err != nil {
		return err
	}
	return nil
}
