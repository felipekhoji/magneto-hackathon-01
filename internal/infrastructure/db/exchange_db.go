package db

import "database/sql"

type ExchangeDB interface {
	AddExchangeRate(fromCurrency string, toCurrency string, rate float64) error
	GetExchangeRate(fromCurrency string, toCurrency string) (float64, error)
}

type DatabaseImpl struct {
	db *sql.DB
}

func NewExchangeDB(db *sql.DB) ExchangeDB {
	return &DatabaseImpl{
		db: db,
	}
}

func (db *DatabaseImpl) AddExchangeRate(fromCurrency string, toCurrency string, rate float64) error {
	panic("implement me")
}

func (db *DatabaseImpl) GetExchangeRate(fromCurrency string, toCurrency string) (float64, error) {
	panic("implement me")
}
