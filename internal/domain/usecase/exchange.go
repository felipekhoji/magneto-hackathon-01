package usecase

import (
	"fmt"
	"magneto-hackathon-01/internal/infrastructure/datasource"
)

type ExchangeUseCase interface {
	GetExchangeRate(fromCurrency string, toCurrency string) (float64, error)
	PostExchangeRate(fromCurrency string, toCurrency string, rate float64) (float64, error)
}

type exchangeUseCaseImpl struct {
	ExchangeDB datasource.ExchangeDB
}

func NewExchangeUseCase(exchangeDb datasource.ExchangeDB) ExchangeUseCase {
	return &exchangeUseCaseImpl{
		ExchangeDB: exchangeDb,
	}
}

func (e *exchangeUseCaseImpl) GetExchangeRate(fromCurrency string, toCurrency string) (float64, error) {
	fmt.Println("[GetExchangeRate]")
	rate, err := e.ExchangeDB.GetExchangeRate(fromCurrency, toCurrency)
	if err != nil {
		return rate, err // TODO: handle error
	}

	return rate, nil
}

func (e *exchangeUseCaseImpl) PostExchangeRate(fromCurrency string, toCurrency string, rate float64) (float64, error) {
	fmt.Println("[PostExchangeRate]")
	err := e.ExchangeDB.AddExchangeRate(fromCurrency, toCurrency, rate)
	if err != nil {
		return rate, err // TODO: handle error
	}

	return rate, nil
}
