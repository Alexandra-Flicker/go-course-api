package repository

import (
	"github.com/jmoiron/sqlx"
	"newProject_courses/internal/entity"
)

type CourseRepo interface {
	GetAll() ([]entity.Course, error)
	Create(course entity.Course) error
	UpdateDescription(course *entity.Course) error
	GetByID(id int) (entity.Course, error)
	DeleteByID(id int) error
}
type courseRepo struct {
	db *sqlx.DB
}

// конструктор
func NewCourseRepo(db *sqlx.DB) CourseRepo {
	return &courseRepo{db}
}

func (r *courseRepo) GetAll() ([]entity.Course, error) {
	var courses []entity.Course
	err := r.db.Select(&courses, "SELECT * FROM courses ORDER BY updated_at DESC")
	return courses, err
}

func (r *courseRepo) Create(course entity.Course) error {
	query := "INSERT INTO courses(name, description, created_at, updated_at) VALUES ($1, $2, now(), now())"
	_, err := r.db.Exec(query, course.Name, course.Description)
	return err
}

func (r *courseRepo) UpdateDescription(course *entity.Course) error {
	query := "UPDATE courses SET description = $1, updated_at = now() WHERE id = $2"
	_, err := r.db.Exec(query, course.Description, course.ID)
	return err
}

func (r *courseRepo) GetByID(id int) (entity.Course, error) {
	var course entity.Course
	err := r.db.Get(&course, "SELECT * FROM courses WHERE id = $1", id)
	return course, err
}

func (r *courseRepo) DeleteByID(id int) error {
	_, err := r.db.Exec("DELETE FROM courses WHERE id = $1", id)
	return err
}
