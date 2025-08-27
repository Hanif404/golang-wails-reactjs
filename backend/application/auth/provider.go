package auth

import (
	"database/sql"
	authhandler "golang-wails-reactjs/backend/application/auth/handler"
	authrepository "golang-wails-reactjs/backend/application/auth/repository"
	authservice "golang-wails-reactjs/backend/application/auth/service"
	"golang-wails-reactjs/backend/domain/auth"
)

type App struct {
	db      *sql.DB
	handler *authhandler.Handler
}

func Provider(db *sql.DB) *App {
	authRepo := authrepository.New(db)
	authService := authservice.New(authRepo)
	authHandler := authhandler.New(authService)

	return &App{
		db:      db,
		handler: authHandler,
	}
}

func (a *App) ActionLogin(payload auth.Login) string {
	return a.handler.PostLogin(payload)
}
