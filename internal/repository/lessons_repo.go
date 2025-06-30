package repository

import (
	"github.com/jmoiron/sqlx"
	"newProject_courses/internal/entity"
)

type LessonsRepo interface {
	CreateLesson(lesson *entity.Lesson) error
	GetLessonByID(id int) (entity.Lesson, error)
	GetAllLessons() ([]entity.Lesson, error)
	UpdateLessonByID(id int, lesson *entity.Lesson) error
	DeleteLessonByID(id int) error
}

type lessonsRepo struct {
	db *sqlx.DB
}

func NewLessonRepo(db *sqlx.DB) LessonsRepo {
	return &lessonsRepo{db: db}
}

func (l *lessonsRepo) CreateLesson(lesson *entity.Lesson) error {
	query := "INSERT INTO lessons (module_id, title, content, created_at, updated_at) VALUES ($1, $2, $3, now(), now())"
	_, err := l.db.Exec(query, lesson.ModuleID, lesson.Title, lesson.Content)
	return err
}

func (l *lessonsRepo) GetLessonByID(id int) (entity.Lesson, error) {
	var lesson entity.Lesson
	query := "SELECT * FROM lessons WHERE id=$1"
	err := l.db.Get(&lesson, query, id)
	return lesson, err
}

func (l *lessonsRepo) GetAllLessons() ([]entity.Lesson, error) {
	var lessons []entity.Lesson
	query := "SELECT * FROM lessons"
	err := l.db.Select(&lessons, query)
	return lessons, err

}

func (l *lessonsRepo) UpdateLessonByID(id int, lesson *entity.Lesson) error {
	query := "UPDATE lessons SET module_id=$1, title=$2, content=$3, updated_at=now() WHERE id=$4"
	_, err := l.db.Exec(query, lesson.ModuleID, lesson.Title, lesson.Content, id)
	return err
}

func (l *lessonsRepo) DeleteLessonByID(id int) error {
	query := "DELETE FROM lessons WHERE id=$1"
	_, err := l.db.Exec(query, id)
	return err
}
