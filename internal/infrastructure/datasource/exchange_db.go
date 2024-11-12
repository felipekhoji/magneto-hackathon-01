package datasource

import (
	"magneto-hackathon-01/pkg/database"
)

// ExchangeDB defines an interface for interacting with exchange database.
// It supports retrieving and storing data through Getter and Poster interfaces.
type ExchangeDB interface {
	Getter
	Poster
}

type ExchangeDBImpl struct {
}

func NewExchangeDB() ExchangeDB {
	return &ExchangeDBImpl{}
}

func (db *ExchangeDBImpl) AddExchangeRate(fromCurrency string, toCurrency string, rate float64) error {
	err := database.AddExchangeRate(fromCurrency, toCurrency, rate)
	if err != nil {
		return err // TODO: handle error
	}

	return nil
}

func (db *ExchangeDBImpl) GetExchangeRate(fromCurrency string, toCurrency string) (float64, error) {
	rate, err := database.GetExchangeRate(fromCurrency, toCurrency)
	if err != nil {
		return rate, err // TODO: handle error
	}

	return rate, nil
}
