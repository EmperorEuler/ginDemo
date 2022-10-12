// +ignore
package routers

import (
	"github.com/gin-gonic/gin"
	"goDemo/controller"
)

func BeforeAuthRouter(r *gin.Engine) *gin.Engine {
	r.POST("/api/auth/register", controller.Register)
	r.POST("/api/auth/login", controller.Login)
	return r
}
