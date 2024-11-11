package usecase

import "fmt"

type ExchangeUseCase interface {
	GetExchangeRate(fromCurrency string, toCurrency string) (float64, error)
	PostExchangeRate(fromCurrency string, toCurrency string, rate float64) (float64, error)
}

type exchangeUseCaseImpl struct {
}

func NewExchangeUseCase() ExchangeUseCase {
	return &exchangeUseCaseImpl{}
}

func (e *exchangeUseCaseImpl) GetExchangeRate(fromCurrency string, toCurrency string) (float64, error) {
	fmt.Println("[GetExchangeRate]")
	return 0, nil
}

func (e *exchangeUseCaseImpl) PostExchangeRate(fromCurrency string, toCurrency string, rate float64) (float64, error) {
	fmt.Println("[PostExchangeRate]")
	return 0, nil
}
