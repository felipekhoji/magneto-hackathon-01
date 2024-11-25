package middleware

import (
	"fmt"
	"github.com/aidarkhanov/nanoid"
	"github.com/gin-gonic/gin"
)

// TODO: is it in the right place?
// AddRequestId adds a request id to the request header
func AddRequestId(c *gin.Context) {
	fmt.Println("[middleware] - AddRequestId")

	requestId, err := generateRequestId()
	if err != nil {
		fmt.Println("Error generating request id:", err)
	}
	c.Request.Header.Add("request-id", requestId)

	c.Next()
}

func generateRequestId() (string, error) {
	alphabet := nanoid.DefaultAlphabet
	size := nanoid.DefaultSize

	requestId, err := nanoid.Generate(alphabet, size)
	if err != nil {
		return "", err
	}
	return requestId, nil
}
