package app

import (
	"github.com/gin-gonic/gin"
	"magneto-hackathon-01/cmd/api/app/dependency"
)

func RouterMap(router *gin.Engine, c *dependency.Controllers) {
	router.GET("/ping", c.PingController.Ping)
	router.GET("/exchange-rate", c.ExchangeRateController.GetExchangeRate)
	router.POST("/exchange-rate", c.ExchangeRateController.PostExchangeRate)
}
