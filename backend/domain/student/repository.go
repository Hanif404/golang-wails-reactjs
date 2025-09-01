package student

type Repository interface {
	GetStudent(id int) (*Student, error)
	GetStudents() (*[]Student, error)
	CreateStudent(student *Student) error
	UpdateStudent(student *Student) error
	DeleteStudent(id int) error
}
