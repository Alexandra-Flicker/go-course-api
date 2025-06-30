package service

import (
	"newProject_courses/internal/entity"
	"newProject_courses/internal/repository"
)

type LessonService interface {
	CreateLesson(lesson *entity.Lesson) error
	GetLessonByID(id int) (entity.Lesson, error)
	GetAllLessons() ([]entity.Lesson, error)
	UpdateLessonByID(id int, lesson *entity.Lesson) error
	DeleteLessonByID(id int) error
}

type lessonService struct {
	repo repository.LessonsRepo
}

func NewLessonService(r repository.LessonsRepo) LessonService {
	return &lessonService{r}
}

func (l *lessonService) CreateLesson(lesson *entity.Lesson) error {
	return l.repo.CreateLesson(lesson)
}

func (l *lessonService) GetLessonByID(id int) (entity.Lesson, error) {
	return l.repo.GetLessonByID(id)
}

func (l *lessonService) GetAllLessons() ([]entity.Lesson, error) {
	return l.repo.GetAllLessons()
}

func (l *lessonService) UpdateLessonByID(id int, lesson *entity.Lesson) error {
	return l.repo.UpdateLessonByID(id, lesson)
}

func (l *lessonService) DeleteLessonByID(id int) error {
	return l.repo.DeleteLessonByID(id)
}
