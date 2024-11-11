package usecase

import (
	"fmt"
	"magneto-hackathon-01/internal/infrastructure/db"
)

type ExchangeUseCase interface {
	GetExchangeRate(fromCurrency string, toCurrency string) (float64, error)
	PostExchangeRate(fromCurrency string, toCurrency string, rate float64) (float64, error)
}

type exchangeUseCaseImpl struct {
	ExchangeDB db.ExchangeDB
}

func NewExchangeUseCase(exchangeDb db.ExchangeDB) ExchangeUseCase {
	return &exchangeUseCaseImpl{
		ExchangeDB: exchangeDb,
	}
}

func (e *exchangeUseCaseImpl) GetExchangeRate(fromCurrency string, toCurrency string) (float64, error) {
	fmt.Println("[GetExchangeRate]")
	e.ExchangeDB.GetExchangeRate(fromCurrency, toCurrency)
	return 0, nil
}

func (e *exchangeUseCaseImpl) PostExchangeRate(fromCurrency string, toCurrency string, rate float64) (float64, error) {
	fmt.Println("[PostExchangeRate]")
	e.ExchangeDB.AddExchangeRate(fromCurrency, toCurrency, rate)
	return 0, nil
}
