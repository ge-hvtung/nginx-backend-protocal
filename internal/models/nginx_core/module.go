package nginxcore

type CoreParser interface {
	Parse() (CoreModule, error)
}

// Init Nginx Core module
func NewCoreModule() *CoreModule {
	return &CoreModule{}
}
