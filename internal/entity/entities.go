package entity

import "time"

type Course struct {
	ID          int       `db:"id" json:"id"`
	Name        string    `db:"name" json:"name"`
	Description string    `db:"description" json:"description"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time `db:"updated_at" json:"updated_at"`
}

type Module struct {
	ID        int       `db:"id" json:"id"`
	CourseID  int       `db:"course_id" json:"course_id"`
	Title     string    `db:"title" json:"title"`
	CreatedAT time.Time `db:"created_at" json:"created_at"`
	UpdatedAT time.Time `db:"updated_at" json:"updated_at"`
}

type Lesson struct {
	ID        int       `db:"id" json:"id"`
	ModuleID  int       `db:"module_id" json:"module_id"`
	Title     string    `db:"title" json:"title"`
	Content   string    `db:"content" json:"content"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}
