package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"goDemo/common"
	"goDemo/middleware"
	"goDemo/routers"
	"os"
)

func main() {
	InitConfig()
	common.InitDB()
	defer common.DB.Close()
	r := gin.Default()
	r.Use(middleware.Cors(), middleware.Recovery())
	r = routers.BeforeAuthRouter(r)
	r.Use(middleware.Auth())
	r = routers.AuthRouter(r)
	port := viper.GetString("server.port")
	if port != "" {
		panic(r.Run(":" + port))
	}
	panic(r.Run())
}

func InitConfig() {
	wordDir, _ := os.Getwd()
	viper.SetConfigName("application.yml")
	viper.SetConfigType("yml")
	viper.AddConfigPath(wordDir + "/config")
	err := viper.ReadInConfig()
	if err != nil {
		panic("config err:" + err.Error())
	}
}
