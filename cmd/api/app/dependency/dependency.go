package dependency

type Dependency struct {
	Gateway    *Gateways
	UseCase    *Usecases
	Controller *Controllers
}

func ResolveDependencies() *Dependency {
	dependencies := &Dependency{}

	return dependencies.
		newGateway().
		newUseCase().
		newController()
}
