package service

import (
	"newProject_courses/internal/entity"
	"newProject_courses/internal/repository"
)

type ModuleService interface {
	CreateModule(module entity.Module) error
	GetAllModules() ([]entity.Module, error)
	GetModuleByID(id int) (entity.Module, error)
	UpdateModuleTitleByID(id int, module entity.Module) error
	DeleteModuleByID(id int) error
}

type moduleService struct {
	repo repository.ModulesRepo
}

// Конструктор
func NewModuleService(r repository.ModulesRepo) ModuleService {
	return &moduleService{r}
}

func (s *moduleService) CreateModule(module entity.Module) error {
	return s.repo.CreateModule(module)
}

func (s *moduleService) GetAllModules() ([]entity.Module, error) {
	return s.repo.GetAllModules()
}

func (s *moduleService) GetModuleByID(id int) (entity.Module, error) {
	return s.repo.GetModuleByID(id)
}

func (s *moduleService) UpdateModuleTitleByID(id int, module entity.Module) error {
	return s.repo.UpdateModuleTitleByID(id, module)
}

func (s *moduleService) DeleteModuleByID(id int) error {
	return s.repo.DeleteModuleByID(id)
}
