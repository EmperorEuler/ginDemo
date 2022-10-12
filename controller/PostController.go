package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"goDemo/common"
	"goDemo/model"
	"goDemo/response"
	"strconv"
)

type IPostController interface {
	RestController
	PageList(c *gin.Context)
}

type PostController struct {
	DB *gorm.DB
}

func (p PostController) PageList(c *gin.Context) {
	//获取分页参数
	pageNum, _ := strconv.Atoi(c.DefaultQuery("pageNum", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "20"))
	//分页
	var posts []model.Post
	p.DB.Order("created_at desc").Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(&posts)
	//前端渲染分页需要知道总数
	var total int
	p.DB.Model(model.Post{}).Count(&total)

	response.Success(c, "成功", gin.H{"data": posts, "total": total})
}

func (p PostController) Create(c *gin.Context) {
	var requestPost model.Post
	//数据验证
	if err := c.ShouldBind(&requestPost); err != nil {
		panic(err)
		return
	}
	//获取登录用户 user
	user, _ := c.Get("user")
	//创建post
	post := model.Post{
		UserId:     user.(model.User).ID,
		CategoryId: requestPost.CategoryId,
		Category:   requestPost.Category,
		Title:      requestPost.Title,
		HeadImg:    requestPost.HeadImg,
		Content:    requestPost.Content,
	}
	//插入数据
	if err := p.DB.Create(&post).Error; err != nil {
		panic(err)
		return
	}
	response.Success(c, "创建成功", nil)
}

func (p PostController) Update(c *gin.Context) {
	var requestPost model.Post
	//数据验证
	if err := c.ShouldBind(&requestPost); err != nil {
		panic(err)
		return
	}

	//获取path中的id
	postId := c.Params.ByName("id")

	var post model.Post
	if p.DB.Where("id = ?", postId).First(&post).RecordNotFound() {
		response.Fail(c, "文章不存在", nil)
		return
	}
	//判断当前用户是否为文章作者
	//获取当前登录用户
	user, _ := c.Get("user")
	userId := user.(model.User).ID
	if userId != post.UserId {
		response.Fail(c, "文章不属于您,请勿非法操作", nil)
		return
	}
	//更新文章
	if err := p.DB.Model(&post).Update(requestPost).Error; err != nil {
		response.Fail(c, "修改失败", nil)
		return
	}

	response.Success(c, "修改成功", gin.H{"post": post})
}

func (p PostController) Show(c *gin.Context) {
	//获取path中的id
	postId := c.Params.ByName("id")

	var post model.Post
	if p.DB.Preload("Category").Where("id = ?", postId).First(&post).RecordNotFound() {
		response.Fail(c, "文章不存在", nil)
		return
	}

	response.Success(c, "", gin.H{"post": post})
}

func (p PostController) Delete(c *gin.Context) {
	//获取path中的id
	postId := c.Params.ByName("id")

	var post model.Post
	if p.DB.Where("id = ?", postId).First(&post).RecordNotFound() {
		response.Fail(c, "文章不存在", nil)
		return
	}
	//判断当前用户是否为文章作者
	//获取当前登录用户
	user, _ := c.Get("user")
	userId := user.(model.User).ID
	if userId != post.UserId {
		response.Fail(c, "文章不属于您,请勿非法操作", nil)
		return
	}

	p.DB.Delete(&post)

	response.Success(c, "删除成功", nil)

}

func NewPostController() IPostController {
	common.DB.AutoMigrate(model.Post{})
	return PostController{DB: common.DB}
}
