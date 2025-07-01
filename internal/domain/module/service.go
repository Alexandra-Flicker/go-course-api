package module

type Service interface {
	CreateModule(module Module) error
	GetAllModules() ([]Module, error)
	GetModuleByID(id int) (Module, error)
	UpdateModuleTitleByID(id int, module Module) error
	DeleteModuleByID(id int) error
}

type moduleService struct {
	repo Repository
}

// NewModuleService конструктор
func NewModuleService(r Repository) Service {
	return &moduleService{r}
}

func (s *moduleService) CreateModule(module Module) error {
	return s.repo.CreateModule(module)
}

func (s *moduleService) GetAllModules() ([]Module, error) {
	return s.repo.GetAllModules()
}

func (s *moduleService) GetModuleByID(id int) (Module, error) {
	return s.repo.GetModuleByID(id)
}

func (s *moduleService) UpdateModuleTitleByID(id int, module Module) error {
	return s.repo.UpdateModuleTitleByID(id, module)
}

func (s *moduleService) DeleteModuleByID(id int) error {
	return s.repo.DeleteModuleByID(id)
}
