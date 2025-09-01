package student

type Service interface {
	GetStudent(id int) (*Student, error)
	GetStudents() (*[]Student, error)
	SaveData(student *Student) error
	DeleteData(id int) error
}
