// +ignore
package routers

import (
	"github.com/gin-gonic/gin"
	"goDemo/controller"
)

func AuthRouter(r *gin.Engine) *gin.Engine {
	r.GET("api/auth/info", controller.Info)
	categoryRouters := r.Group("/categorys")
	categoryController := controller.NewCategoryController()
	categoryRouters.POST("", categoryController.Create)
	categoryRouters.PUT("/:id", categoryController.Update)
	categoryRouters.GET("/:id", categoryController.Show)
	categoryRouters.DELETE("/:id", categoryController.Delete)

	postRouters := r.Group("/posts")
	postController := controller.NewPostController()
	postRouters.POST("", postController.Create)
	postRouters.PUT("/:id", postController.Update)
	postRouters.GET("/:id", postController.Show)
	postRouters.DELETE("/:id", postController.Delete)
	postRouters.POST("/page/list", postController.PageList)
	return r
}
