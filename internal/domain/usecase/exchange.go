package usecase

import (
	"fmt"
	"magneto-hackathon-01/internal/domain/entity"
	"magneto-hackathon-01/internal/domain/gateway"
)

type ExchangeUseCase interface {
	GetExchangeRate(fromCurrency string, toCurrency string) (*entity.Exchange, error)
	PostExchangeRate(fromCurrency string, toCurrency string, rate float64) (*entity.Exchange, error)
}

type exchangeUseCaseImpl struct {
	ExchangeRateGateway gateway.ExchangeRateGateway
}

func NewExchangeUseCase(gateway gateway.ExchangeRateGateway) ExchangeUseCase {
	return &exchangeUseCaseImpl{
		ExchangeRateGateway: gateway,
	}
}

func (e *exchangeUseCaseImpl) GetExchangeRate(fromCurrency string, toCurrency string) (*entity.Exchange, error) {
	fmt.Println("[usecase] - GetExchangeRate")
	rate, err := e.ExchangeRateGateway.GetExchangeRate(fromCurrency, toCurrency)
	if err != nil {
		return nil, fmt.Errorf("not found") // TODO: handle error
	}

	return &entity.Exchange{
		FromCurrency: fromCurrency,
		ToCurrency:   toCurrency,
		Rate:         rate,
	}, nil
}

func (e *exchangeUseCaseImpl) PostExchangeRate(fromCurrency string, toCurrency string, rate float64) (*entity.Exchange, error) {
	fmt.Println("[usecase] - PostExchangeRate")
	err := e.ExchangeRateGateway.AddExchangeRate(fromCurrency, toCurrency, rate)
	if err != nil {
		return nil, err // TODO: handle error
	}

	return &entity.Exchange{
		FromCurrency: fromCurrency,
		ToCurrency:   toCurrency,
		Rate:         rate,
	}, nil
}
