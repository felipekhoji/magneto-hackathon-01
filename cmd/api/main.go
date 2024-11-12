package main

import (
	"github.com/gin-gonic/gin"
	"magneto-hackathon-01/cmd/config"
	"magneto-hackathon-01/internal/domain/usecase"
	"magneto-hackathon-01/internal/infrastructure/db"
	"magneto-hackathon-01/pkg/database"
	"os"
	"strconv"
)

func main() {

	// set up database
	absPath := os.Getenv("PWD")
	databaseFilePath := absPath + config.LocalDBPath
	err := database.InitDB(databaseFilePath)
	if err != nil {
		panic(err) // TODO: handle error
	}

	// inject dependencies
	exchangeDB := db.NewExchangeDB()
	exchange := usecase.NewExchangeUseCase(exchangeDB)

	// ref. https://gin-gonic.com/docs/quickstart/#getting-started
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
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

		c.JSON(200, gin.H{
			"rate": exchangeRate,
		})
	})

	r.POST("/exchange-rate", func(c *gin.Context) {
		fromCurrency := c.Query("from")
		toCurrency := c.Query("to")
		exchangeRate, _ := strconv.ParseFloat(c.Query("rate"), 64)
		exchangeRateResp, err := exchange.PostExchangeRate(fromCurrency, toCurrency, exchangeRate) // TODO: query params
		if err != nil {
			// TODO: handle error
			c.JSON(500, gin.H{
				"message": "error posting exchange rate",
			})
			return
		}

		c.JSON(200, gin.H{
			"converted_amount": exchangeRateResp,
		})
	})

	r.Run() // listen and serve on 0.0.0.0:8080
}
