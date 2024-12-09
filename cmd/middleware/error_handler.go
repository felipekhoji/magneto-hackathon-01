package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

// TODO: is it in the right place?
func ErrorHandler(c *gin.Context) {
	c.Next()

	lastErr := c.Errors.Last()
	if lastErr == nil {
		fmt.Println("[middleware] - ErrorHandler - no errors to handle")
		return
	}

	// TODO: handle error for each case, 500, 400, etc.
	fmt.Println("[middleware] - ErrorHandler - error to handle:", lastErr.Error())
	if lastErr.Error() == "taxa de câmbio não encontrada" {
		c.JSON(404, gin.H{
			"message": "taxa de câmbio não encontrada",
		})
	} else {
		c.JSON(500, gin.H{
			"message": "internal server error",
		})
	}

	c.Abort() // TODO: abort to prevent execution of other handlers, but how to use it?

}
