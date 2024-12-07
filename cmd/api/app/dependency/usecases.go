package dependency

import "magneto-hackathon-01/internal/domain/usecase"

type Usecases struct {
	ExchangeUseCase usecase.ExchangeUseCase
}

func (d *Dependency) newUseCase() *Dependency {
	d.UseCase = &Usecases{
		ExchangeUseCase: usecase.NewExchangeUseCase(d.Gateway.ExchangeRateGateway),
	}
	return d
}
