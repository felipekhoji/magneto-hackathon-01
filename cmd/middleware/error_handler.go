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

	c.Abort() // TODO: abort to prevent execution of other handlers, but how to use it?

}
