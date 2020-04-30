package main

import (
	"fmt"
	"ginp/api"
	"ginp/base"
	"ginp/redispool"
	"ginp/xrom_mysql"
	"github.com/gin-gonic/gin"
	"net/http"
     _ "net/http/pprof"
)

func init()  {
	base.Config()
	xrom_mysql.InitXorm()
	redispool.InitRedis()
}

func main() {
	r := gin.Default()
	r.Use(Cors())
	r.Static("/static", "./static")
	r.GET("/tt", log(), TT)
	r.GET("/model/list", log(), Test)
	apir := r.Group("/api")
	apir.GET("/article/list", api.ArticleList)
	apir.GET("/article/detail", api.ArticleDetail)
	apir.GET("/reply/list", api.ReplyList)
	apir.POST("/reply/save", api.ReplySave)
	apir.GET("/user/login", api.Login)
	apir.POST("/user/save", api.UserSave)
	go func() {
		http.ListenAndServe("localhost:6060", nil)
	}()
	r.Run(":8000")
}

func Test(context *gin.Context)  {
	fmt.Println("Authorization",context.GetHeader("Authorization"))
	context.JSON(200, "ok")
}


func TT(context *gin.Context) {
	s := struct {
		Name string `json:"name"`
	}{Name: "xiaozhang"}
	fmt.Println("222222222222")
	context.JSON(200, s)
}

func log() gin.HandlerFunc {
	return func(context *gin.Context) {
		fmt.Println("ttttttt")
		context.Next()
		fmt.Println("111111111")
	}
}

// 处理跨域请求,支持options访问
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")

		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		// 处理请求
		c.Next()
	}
}
