package lesson

type Service interface {
	CreateLesson(lesson *Lesson) error
	GetLessonByID(id int) (Lesson, error)
	GetAllLessons() ([]Lesson, error)
	UpdateLessonByID(id int, lesson *Lesson) error
	DeleteLessonByID(id int) error
}

type lessonService struct {
	repo Repository
}

func NewLessonService(r Repository) Service {
	return &lessonService{r}
}

func (l *lessonService) CreateLesson(lesson *Lesson) error {
	return l.repo.CreateLesson(lesson)
}

func (l *lessonService) GetLessonByID(id int) (Lesson, error) {
	return l.repo.GetLessonByID(id)
}

func (l *lessonService) GetAllLessons() ([]Lesson, error) {
	return l.repo.GetAllLessons()
}

func (l *lessonService) UpdateLessonByID(id int, lesson *Lesson) error {
	return l.repo.UpdateLessonByID(id, lesson)
}

func (l *lessonService) DeleteLessonByID(id int) error {
	return l.repo.DeleteLessonByID(id)
}
