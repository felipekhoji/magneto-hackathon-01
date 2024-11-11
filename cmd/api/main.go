package main

import (
	"github.com/gin-gonic/gin"
	"magneto-hackathon-01/cmd/config"
	"magneto-hackathon-01/internal/domain/usecase"
	"magneto-hackathon-01/internal/infrastructure/db"
	"magneto-hackathon-01/pkg/database"
	"os"
)

func main() {
	absPath := os.Getenv("PWD")
	databaseFilePath := absPath + config.LocalDBPath
	err := database.InitDB(databaseFilePath)
	if err != nil {
		panic(err) // TODO: handle error
	}

	exchangeDB := db.NewExchangeDB(nil)
	exchange := usecase.NewExchangeUseCase(exchangeDB)

	// ref. https://gin-gonic.com/docs/quickstart/#getting-started
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/exchange-rate", func(c *gin.Context) {
		exchangeRate, err := exchange.GetExchangeRate("BRL", "USD")
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
		exchangeRate, err := exchange.PostExchangeRate("BRL", "USD", 6.00)
		if err != nil {
			// TODO: handle error
			c.JSON(500, gin.H{
				"message": "error posting exchange rate",
			})
			return
		}

		c.JSON(200, gin.H{
			"converted_amount": exchangeRate,
		})
	})

	r.Run() // listen and serve on 0.0.0.0:8080
}
