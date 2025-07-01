package course

type Repository interface {
	GetAll() ([]Course, error)
	Create(course Course) error
	UpdateDescription(course *Course) error
	GetByID(id int) (Course, error)
	DeleteByID(id int) error
}
