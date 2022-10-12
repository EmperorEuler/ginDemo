// +ignore
package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"goDemo/common"
	"goDemo/dto"
	"goDemo/model"
	"goDemo/response"
	"goDemo/util"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
)

func Info(c *gin.Context) {
	user, _ := c.Get("user")
	response.Response(c, http.StatusOK, 200, gin.H{"user": dto.ToUserDto(user.(model.User))}, "")
}

func Register(c *gin.Context) {
	//获取参数
	//使用map获取请求参数
	//var requestMap = make(map[string]string)
	//json.NewDecoder(c.Request.Body).Decode(&requestMap)

	//使用结构体获取请求参数
	var requestUser model.User
	//json.NewDecoder(c.Request.Body).Decode(&requestUser)
	c.Bind(&requestUser)

	name := requestUser.Name
	telephone := requestUser.Telephone
	password := requestUser.Password
	//数据验证
	if len(telephone) != 11 {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "手机号必须为11位")
		return
	}
	if len(password) < 6 {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "密码不少于6位")
		return
	}
	//如果名称没有传，给一个10位的随机字符串
	if len(name) == 0 {
		name = util.RandString(10)
	}
	//判断手机号是否存在
	if util.IsTelephoneExist(common.DB, telephone) {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "用户已经存在")
		return
	}
	//创建用户
	//创建用户前密码加密
	hasedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		response.Response(c, http.StatusInternalServerError, 500, nil, "加密错误")
		return
	}
	newUser := model.User{
		Model:     gorm.Model{},
		Name:      name,
		Telephone: telephone,
		Password:  string(hasedPassword),
	}
	common.DB.Create(&newUser)

	//发放token
	token, err := common.ReleaseToken(newUser)
	if err != nil {
		response.Response(c, http.StatusInternalServerError, 500, nil, "系统异常")
		log.Panicf("token生成失败,错误:%s", err)
	}

	response.Success(c, "注册成功", gin.H{"token": token})
}

func Login(c *gin.Context) {
	//获取数据
	var requestUser model.User
	c.Bind(&requestUser)
	telephone := requestUser.Telephone
	password := requestUser.Password
	//数据验证
	if len(telephone) != 11 {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "手机号必须为11位")
		return
	}
	if len(password) < 6 {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "密码不少于6位")
		return
	}
	//判断手机号是否存在
	var user model.User
	common.DB.Where("telephone = ?", telephone).First(&user)
	if user.ID == 0 {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "用户不存在")
	}
	//判断密码是否正确
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		response.Response(c, http.StatusBadRequest, 400, nil, "密码错误")
		return
	}
	//发放token
	token, err := common.ReleaseToken(user)
	if err != nil {
		response.Response(c, http.StatusInternalServerError, 500, nil, "系统异常")
		log.Panicf("token生成失败,错误:%s", err)
	}
	//返回结果
	response.Success(c, "登录成功", gin.H{"token": token})
}
