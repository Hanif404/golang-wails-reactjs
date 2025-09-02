package studentservice

import (
	"golang-wails-reactjs/backend/domain/student"
	"time"

	"github.com/google/uuid"
)

type service struct {
	repo student.Repository
}

func New(r student.Repository) student.Service {
	return &service{
		repo: r,
	}
}

func (s *service) GetStudents() (*[]student.Student, error) {
	return s.repo.GetStudents()
}

func (s *service) GetStudent(id int) (*student.Student, error) {
	return s.repo.GetStudent(id)
}

func (s *service) SaveData(student *student.Student) error {
	if student.ID == 0 {
		student.CreatedAt = time.Now().UnixMilli()
		student.Key = uuid.New().String()

		return s.repo.CreateStudent(student)
	}
	return s.repo.UpdateStudent(student)
}

func (s *service) DeleteData(id int) error {
	return s.repo.DeleteStudent(id)
}
