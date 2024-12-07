package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"magneto-hackathon-01/cmd/api/controller/dto"
	"magneto-hackathon-01/cmd/config"
	"magneto-hackathon-01/cmd/middleware"
	"magneto-hackathon-01/internal/domain/usecase"
	"magneto-hackathon-01/internal/infrastructure/db"
	"magneto-hackathon-01/pkg/database"
	"os"
	"strconv"
)

// TODO create a treatment for errors, generic and specifics errors

func main() {

	// set up database
	absPath := os.Getenv("PWD")
	databaseFilePath := absPath + config.LocalDBPath
	err := database.InitDB(databaseFilePath)
	if err != nil {
		panic(err) // TODO: handle error
	}

	// inject dependencies
	exchangeDB := db.NewExchangeRateDB()
	exchange := usecase.NewExchangeUseCase(exchangeDB)

	// ref. https://gin-gonic.com/docs/quickstart/#getting-started
	r := gin.Default()
	r.Use(middleware.ErrorHandler)

	r.Use(middleware.AddRequestId)

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/error-handler-test", func(c *gin.Context) {
		if true {
			c.Error(fmt.Errorf("error test", 500))
		}
		c.JSON(500, gin.H{
			"message": "pong",
		})

	})

	r.GET("/exchange-rate", func(c *gin.Context) {
		fromCurrency := c.Query("from")
		toCurrency := c.Query("to")
		exchangeRate, err := exchange.GetExchangeRate(fromCurrency, toCurrency) // TODO: query params
		if err != nil {
			// TODO: handle error
			c.JSON(500, gin.H{
				"message": "error getting exchange rate",
			})
			return
		}

		fmt.Fprintf(os.Stdout, "/exchange-rate - request-id: %s\n", c.GetHeader("request-id"))

		exchangeRateDTO := dto.ExchangeRateDTO{
			Rate: exchangeRate.Rate,
		}
		c.JSON(200, exchangeRateDTO)
	})

	r.POST("/exchange-rate", func(c *gin.Context) {
		fromCurrency := c.Query("from")
		toCurrency := c.Query("to")
		exchangeRate, _ := strconv.ParseFloat(c.Query("rate"), 64)
		exchangeRateResp, err := exchange.PostExchangeRate(fromCurrency, toCurrency, exchangeRate)
		if err != nil {
			// TODO: handle error
			c.JSON(500, gin.H{
				"message": "error posting exchange rate",
			})
			return
		}

		c.JSON(200, exchangeRateResp)
	})

	r.Run() // listen and serve on 0.0.0.0:8080
}
