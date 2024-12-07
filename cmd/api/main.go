package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"magneto-hackathon-01/cmd/api/app"
	"magneto-hackathon-01/cmd/api/app/dependency"
	"magneto-hackathon-01/cmd/config"
	"magneto-hackathon-01/cmd/middleware"
	"magneto-hackathon-01/pkg/database"
	"os"
)

// TODO create a treatment for errors, generic and specifics errors

func main() {

	// TODO: this could be moved to a builder step
	// set up database
	absPath := os.Getenv("PWD")
	databaseFilePath := absPath + config.LocalDBPath
	err := database.InitDB(databaseFilePath)
	if err != nil {
		panic(err) // TODO: handle error
	}

	// inject dependencies
	dependency := dependency.ResolveDependencies()

	// TODO: maybe this could be moved to a builder step too
	// ref. https://gin-gonic.com/docs/quickstart/#getting-started
	r := gin.Default()
	r.Use(middleware.ErrorHandler)
	r.Use(middleware.AddRequestId)
	app.RouterMap(r, dependency.Controller)

	r.GET("/error-handler-test", func(c *gin.Context) {
		if true {
			c.Error(fmt.Errorf("error test", 500))
		}
		c.JSON(500, gin.H{
			"message": "pong",
		})

	})

	r.Run() // listen and serve on 0.0.0.0:8080
}
