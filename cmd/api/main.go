package main

import (
	"github.com/gin-gonic/gin"
	"magneto-hackathon-01/internal/domain/usecase"
)

func main() {

	exchange := usecase.NewExchangeUseCase()

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
