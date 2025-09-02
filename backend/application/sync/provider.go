package sync

import (
	"database/sql"
	studentRepository "golang-wails-reactjs/backend/application/student/repository"
	handler "golang-wails-reactjs/backend/application/sync/handler"
	repository "golang-wails-reactjs/backend/application/sync/repository"
	service "golang-wails-reactjs/backend/application/sync/service"
)

type App struct {
	db      *sql.DB
	handler *handler.Handler
}

func Provider(db *sql.DB) *App {
	repo := repository.New(db)
	studentRepo := studentRepository.New(db)
	service := service.New(repo, studentRepo)
	handler := handler.New(service)

	return &App{
		db:      db,
		handler: handler,
	}
}

func (a *App) ActionSync(email string) string {
	return a.handler.PostData(email)
}

func (a *App) Syncs() string {
	return a.handler.GetDatas()
}

func (a *App) Sync(id int) string {
	return a.handler.GetData(id)
}
