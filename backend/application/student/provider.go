package student

import (
	"database/sql"
	studenthandler "golang-wails-reactjs/backend/application/student/handler"
	studentrepository "golang-wails-reactjs/backend/application/student/repository"
	studentservice "golang-wails-reactjs/backend/application/student/service"
	"golang-wails-reactjs/backend/domain/student"
)

type App struct {
	db      *sql.DB
	handler *studenthandler.Handler
}

func Provider(db *sql.DB) *App {
	repo := studentrepository.New(db)
	service := studentservice.New(repo)
	handler := studenthandler.New(service)

	return &App{
		db:      db,
		handler: handler,
	}
}

func (a *App) ActionForm(payload student.Student) string {
	return a.handler.PostData(payload)
}

func (a *App) Students() string {
	return a.handler.GetDatas()
}

func (a *App) Student(id int) string {
	return a.handler.GetData(id)
}

func (a *App) ActionDelete(id int) string {
	return a.handler.DeleteData(id)
}
