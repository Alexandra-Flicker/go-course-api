package lesson

type Repository interface {
	CreateLesson(lesson *Lesson) error
	GetLessonByID(id int) (Lesson, error)
	GetAllLessons() ([]Lesson, error)
	UpdateLessonByID(id int, lesson *Lesson) error
	DeleteLessonByID(id int) error
}
