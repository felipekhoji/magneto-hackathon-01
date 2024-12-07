package dependency

import "magneto-hackathon-01/internal/infrastructure/db"

type Gateways struct {
	ExchangeRateGateway *db.ExchangeRateDBImpl
}

func (d *Dependency) newGateway() *Dependency {
	d.Gateway = &Gateways{
		ExchangeRateGateway: db.NewExchangeRateDB(),
	}
	return d
}
