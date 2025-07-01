package lesson

import "time"

type Lesson struct {
	ID        int       `db:"id" json:"id"`
	ModuleID  int       `db:"module_id" json:"module_id"`
	Title     string    `db:"title" json:"title"`
	Content   string    `db:"content" json:"content"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}
