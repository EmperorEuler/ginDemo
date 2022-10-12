package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"goDemo/response"
)

func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				response.Fail(c, fmt.Sprint(err), nil)
			}
		}()

		c.Next()
	}
}
