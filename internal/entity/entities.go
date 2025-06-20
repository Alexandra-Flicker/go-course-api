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
	ID        int
	CourseID  int
	Title     string
	CreatedAT time.Time
	UpdatedAT time.Time
}

type Lesson struct {
	ID        int
	ModuleID  int
	Title     string
	Content   string
	CreatedAT time.Time
	UpdatedAT time.Time
}
