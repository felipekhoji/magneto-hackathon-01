package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"magneto-hackathon-01/cmd/api/controller/dto"
	"magneto-hackathon-01/internal/domain/usecase"
	"os"
	"strconv"
)

type ExchangeRateController struct {
	exchangeRateUseCase usecase.ExchangeUseCase
}

func NewExchangeRateController(exchangeRateUseCase usecase.ExchangeUseCase) *ExchangeRateController {
	return &ExchangeRateController{
		exchangeRateUseCase: exchangeRateUseCase,
	}
}

func (e ExchangeRateController) GetExchangeRate(c *gin.Context) {
	fromCurrency := c.Query("from")
	toCurrency := c.Query("to")
	exchangeRate, err := e.exchangeRateUseCase.GetExchangeRate(fromCurrency, toCurrency) // TODO: query params
	if err != nil {
		c.Error(err)
		return
	}

	fmt.Fprintf(os.Stdout, "/exchange-rate - request-id: %s\n", c.GetHeader("request-id"))

	exchangeRateDTO := dto.ExchangeRateDTO{
		Rate: exchangeRate.Rate,
	}
	c.JSON(200, exchangeRateDTO)
}

func (e ExchangeRateController) PostExchangeRate(c *gin.Context) {
	fromCurrency := c.Query("from")
	toCurrency := c.Query("to")
	exchangeRate, _ := strconv.ParseFloat(c.Query("rate"), 64)
	exchangeRateResp, err := e.exchangeRateUseCase.PostExchangeRate(fromCurrency, toCurrency, exchangeRate)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(200, exchangeRateResp)
}
