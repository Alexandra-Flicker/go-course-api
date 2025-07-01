package module

type Repository interface {
	CreateModule(module Module) error
	GetAllModules() ([]Module, error)
	GetModuleByID(id int) (Module, error)
	UpdateModuleTitleByID(id int, module Module) error
	DeleteModuleByID(id int) error
}
