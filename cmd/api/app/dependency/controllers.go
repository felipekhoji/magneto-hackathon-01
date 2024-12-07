package dependency

import "magneto-hackathon-01/cmd/api/controller"

type Controllers struct {
	PingController         *controller.PingController
	ExchangeRateController *controller.ExchangeRateController
}

func (d *Dependency) newController() *Dependency {
	d.Controller = &Controllers{
		PingController:         controller.NewPingController(),
		ExchangeRateController: controller.NewExchangeRateController(d.UseCase.ExchangeUseCase),
	}
	return d
}
