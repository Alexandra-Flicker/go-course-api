package module

import (
	"github.com/jmoiron/sqlx"
	"newProject_courses/internal/domain/module"
)

type modulesRepo struct {
	db *sqlx.DB
}

func NewModuleRepo(db *sqlx.DB) module.Repository {
	return &modulesRepo{db: db}
}

func (m *modulesRepo) CreateModule(module module.Module) error {
	query := "INSERT INTO modules (course_id, title, created_at, updated_at) VALUES ($1, $2, now(), now())"
	_, err := m.db.Exec(query, module.CourseID, module.Title, module.CreatedAT, module.UpdatedAT)
	return err
}

func (m *modulesRepo) GetAllModules() ([]module.Module, error) {
	var modules []module.Module
	err := m.db.Select(&modules, "SELECT * FROM modules ORDER BY updated_at DESC")
	return modules, err
}

func (m *modulesRepo) GetModuleByID(id int) (module.Module, error) {
	var moduleResponse module.Module
	err := m.db.Get(&moduleResponse, "SELECT * FROM modules WHERE id=$1", id)
	return moduleResponse, err
}

func (m *modulesRepo) UpdateModuleTitleByID(id int, module module.Module) error {
	_, err := m.db.Exec("UPDATE modules SET title = $1 WHERE id = $2", module.Title, id)
	return err
}

func (m *modulesRepo) DeleteModuleByID(id int) error {
	_, err := m.db.Exec("DELETE FROM modules WHERE id = $1", id)
	return err
}
