package usecase

import (
	"fmt"
	"magneto-hackathon-01/internal/domain/entity"
	"magneto-hackathon-01/internal/infrastructure/datasource"
)

type ExchangeUseCase interface {
	GetExchangeRate(fromCurrency string, toCurrency string) (*entity.Exchange, error)
	PostExchangeRate(fromCurrency string, toCurrency string, rate float64) (*entity.Exchange, error)
}

type exchangeUseCaseImpl struct {
	ExchangeDB datasource.ExchangeDB
}

func NewExchangeUseCase(exchangeDb datasource.ExchangeDB) ExchangeUseCase {
	return &exchangeUseCaseImpl{
		ExchangeDB: exchangeDb,
	}
}

func (e *exchangeUseCaseImpl) GetExchangeRate(fromCurrency string, toCurrency string) (*entity.Exchange, error) {
	fmt.Println("[usecase] - GetExchangeRate")
	rate, err := e.ExchangeDB.GetExchangeRate(fromCurrency, toCurrency)
	if err != nil {
		return nil, err // TODO: handle error
	}

	return &entity.Exchange{
		FromCurrency: fromCurrency,
		ToCurrency:   toCurrency,
		Rate:         rate,
	}, nil
}

func (e *exchangeUseCaseImpl) PostExchangeRate(fromCurrency string, toCurrency string, rate float64) (*entity.Exchange, error) {
	fmt.Println("[usecase] - PostExchangeRate")
	err := e.ExchangeDB.AddExchangeRate(fromCurrency, toCurrency, rate)
	if err != nil {
		return nil, err // TODO: handle error
	}

	return &entity.Exchange{
		FromCurrency: fromCurrency,
		ToCurrency:   toCurrency,
		Rate:         rate,
	}, nil
}
