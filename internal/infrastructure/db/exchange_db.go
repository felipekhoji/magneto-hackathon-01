package db

import (
	"magneto-hackathon-01/pkg/database"
)

// ExchangeRateDBImpl implements the ExchangeRateGateway interface.
type ExchangeRateDBImpl struct {
}

func NewExchangeRateDB() *ExchangeRateDBImpl {
	return &ExchangeRateDBImpl{}
}

func (db *ExchangeRateDBImpl) AddExchangeRate(fromCurrency string, toCurrency string, rate float64) error {
	err := database.AddExchangeRate(fromCurrency, toCurrency, rate)
	if err != nil {
		return err
	}

	return nil
}

func (db *ExchangeRateDBImpl) GetExchangeRate(fromCurrency string, toCurrency string) (float64, error) {
	rate, err := database.GetExchangeRate(fromCurrency, toCurrency)
	if err != nil {
		return rate, err
	}

	return rate, nil
}
