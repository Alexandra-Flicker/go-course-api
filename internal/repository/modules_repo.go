package repository

import (
	"github.com/jmoiron/sqlx"
	"newProject_courses/internal/entity"
)

type ModulesRepo interface {
	CreateModule(module entity.Module) error
	GetAllModules() ([]entity.Module, error)
	GetModuleByID(id int) (entity.Module, error)
	UpdateModuleTitleByID(id int, module entity.Module) error
	DeleteModuleByID(id int) error
}

type modulesRepo struct {
	db *sqlx.DB
}

func NewModuleRepo(db *sqlx.DB) ModulesRepo {
	return &modulesRepo{db: db}
}

func (m *modulesRepo) CreateModule(module entity.Module) error {
	query := "INSERT INTO modules (course_id, title, created_at, updated_at) VALUES ($1, $2, now(), now())"
	_, err := m.db.Exec(query, module.CourseID, module.Title)
	return err
}

func (m *modulesRepo) GetAllModules() ([]entity.Module, error) {
	var modules []entity.Module
	err := m.db.Select(&modules, "SELECT * FROM modules ORDER BY updated_at DESC")
	return modules, err
}

func (m *modulesRepo) GetModuleByID(id int) (entity.Module, error) {
	var module entity.Module
	err := m.db.Get(&module, "SELECT * FROM modules WHERE id=$1", id)
	return module, err
}

func (m *modulesRepo) UpdateModuleTitleByID(id int, module entity.Module) error {
	_, err := m.db.Exec("UPDATE modules SET title = $1 WHERE id = $2", module.Title, id)
	return err
}

func (m *modulesRepo) DeleteModuleByID(id int) error {
	_, err := m.db.Exec("DELETE FROM modules WHERE id = $1", id)
	return err
}
