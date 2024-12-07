package gateway

type ExchangeRateGateway interface {
	AddExchangeRate(fromCurrency string, toCurrency string, rate float64) error
	GetExchangeRate(fromCurrency string, toCurrency string) (float64, error)
}
