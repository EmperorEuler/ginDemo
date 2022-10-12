// +ignore
package controller

import (
	"github.com/gin-gonic/gin"
	"goDemo/model"
	"goDemo/repository"
	"goDemo/response"
	"goDemo/vo"
	"strconv"
)

type ICategoryController interface {
	RestController
}
type CategoryController struct {
	Repository repository.CategoryRepository
}

func NewCategoryController() ICategoryController {
	repository := repository.NewCategoryRepository()
	repository.DB.AutoMigrate(model.Category{})
	return CategoryController{Repository: repository}
}

func (cate CategoryController) Create(c *gin.Context) {
	var requestCategory vo.CreateCategoryRequest
	if err := c.ShouldBind(&requestCategory); err != nil {
		panic(err)
		return
	}
	category, err := cate.Repository.Create(requestCategory.Name)
	if err != nil {
		panic(err)
		return
	}

	response.Success(c, "", gin.H{"category": category})
}

func (cate CategoryController) Update(c *gin.Context) {
	//绑定body中的参数
	var requestCategory vo.CreateCategoryRequest
	if err := c.ShouldBind(&requestCategory); err != nil {
		panic(err)
		return
	}
	//获取path中的参数
	categoryId, _ := strconv.Atoi(c.Params.ByName("id"))

	category, err := cate.Repository.SelectById(categoryId)
	if err != nil {
		panic(err)
		return
	}

	//更新分类
	update, err := cate.Repository.Update(*category, requestCategory.Name)
	if err != nil {
		panic(err)
		return
	}

	response.Success(c, "修改成功", gin.H{"category": update})
}

func (cate CategoryController) Show(c *gin.Context) {
	//获取path中参数
	categoryId, _ := strconv.Atoi(c.Params.ByName("id"))

	category, err := cate.Repository.SelectById(categoryId)
	if err != nil {
		panic(err)
		return
	}

	response.Success(c, "", gin.H{"category": category})
}

func (cate CategoryController) Delete(c *gin.Context) {
	//获取path中参数
	categoryId, _ := strconv.Atoi(c.Params.ByName("id"))

	err := cate.Repository.DeleteById(categoryId)
	if err != nil {
		panic(err)
		return
	}

	response.Success(c, "删除成功", nil)

}
