package course

import (
	"github.com/jmoiron/sqlx"
	"newProject_courses/internal/domain/course"
)

type courseRepo struct {
	db *sqlx.DB
}

// NewCourseRepo constructor
func NewCourseRepo(db *sqlx.DB) course.Repository {
	return &courseRepo{db}
}

func (r *courseRepo) GetAll() ([]course.Course, error) {
	var courses []course.Course
	err := r.db.Select(&courses, "SELECT * FROM courses ORDER BY updated_at DESC")
	return courses, err
}

func (r *courseRepo) Create(course course.Course) error {
	query := "INSERT INTO courses(name, description, created_at, updated_at) VALUES ($1, $2, now(), now())"
	_, err := r.db.Exec(query, course.Name, course.Description)
	return err
}

func (r *courseRepo) UpdateDescription(course *course.Course) error {
	query := "UPDATE courses SET description = $1, updated_at = now() WHERE id = $2"
	_, err := r.db.Exec(query, course.Description, course.ID)
	return err
}

func (r *courseRepo) GetByID(id int) (course.Course, error) {
	var courseResponse course.Course
	err := r.db.Get(&courseResponse, "SELECT * FROM courses WHERE id = $1", id)
	return courseResponse, err
}

func (r *courseRepo) DeleteByID(id int) error {
	_, err := r.db.Exec("DELETE FROM courses WHERE id = $1", id)
	return err
}
