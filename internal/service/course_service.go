package service

import (
	"newProject_courses/internal/entity"
	"newProject_courses/internal/repository"
)

type CourseService interface {
	GetAll() ([]entity.Course, error)
	GetByID(id int) (entity.Course, error)
	CreateCourse(name, description string) error
	UpdateDescriptionByID(description string, id int) error
	DeleteByID(id int) error
}
type courseService struct {
	repo repository.CourseRepo
}

func NewCourseService(r repository.CourseRepo) CourseService {
	return &courseService{r}
}

func (s *courseService) GetAll() ([]entity.Course, error) {
	return s.repo.GetAll()
}

func (s *courseService) CreateCourse(name, description string) error {
	c := &entity.Course{Name: name, Description: description}
	return s.repo.Create(*c)
}
func (s *courseService) UpdateDescriptionByID(description string, ID int) error {
	c := &entity.Course{ID: ID, Description: description}
	return s.repo.UpdateDescription(c)
}
func (s *courseService) GetByID(id int) (entity.Course, error) {
	return s.repo.GetByID(id)
}

func (s *courseService) DeleteByID(id int) error {
	c := entity.Course{ID: id}
	return s.repo.DeleteByID(c.ID)
}
