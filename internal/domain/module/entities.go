package module

import "time"

type Module struct {
	ID        int       `db:"id" json:"id"`
	CourseID  int       `db:"course_id" json:"course_id"`
	Title     string    `db:"title" json:"title"`
	CreatedAT time.Time `db:"created_at" json:"created_at"`
	UpdatedAT time.Time `db:"updated_at" json:"updated_at"`
}
