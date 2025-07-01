package course

type Service interface {
	GetAll() ([]Course, error)
	GetByID(id int) (Course, error)
	CreateCourse(name, description string) error
	UpdateDescriptionByID(description string, id int) error
	DeleteByID(id int) error
}
type courseService struct {
	repo Repository
}

func NewCourseService(r Repository) Service {
	return &courseService{r}
}

func (s *courseService) GetAll() ([]Course, error) {
	return s.repo.GetAll()
}

func (s *courseService) CreateCourse(name, description string) error {
	c := &Course{Name: name, Description: description}
	return s.repo.Create(*c)
}
func (s *courseService) UpdateDescriptionByID(description string, ID int) error {
	c := &Course{ID: ID, Description: description}
	return s.repo.UpdateDescription(c)
}
func (s *courseService) GetByID(id int) (Course, error) {
	return s.repo.GetByID(id)
}

func (s *courseService) DeleteByID(id int) error {
	c := Course{ID: id}
	return s.repo.DeleteByID(c.ID)
}
