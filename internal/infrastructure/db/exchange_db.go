package db

import (
	"magneto-hackathon-01/pkg/database"
)

type ExchangeDB interface {
	AddExchangeRate(fromCurrency string, toCurrency string, rate float64) error
	GetExchangeRate(fromCurrency string, toCurrency string) (float64, error)
}

type ExchangeDBImpl struct {
}

func NewExchangeDB() ExchangeDB {
	return &ExchangeDBImpl{}
}

func (db *DatabaseImpl) AddExchangeRate(fromCurrency string, toCurrency string, rate float64) error {
	panic("implement me")
}

func (db *DatabaseImpl) GetExchangeRate(fromCurrency string, toCurrency string) (float64, error) {
	panic("implement me")
}
