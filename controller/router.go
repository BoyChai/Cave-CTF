package controller

import "github.com/gin-gonic/gin"

// Router 实例化router类型对象，首字母大写用于跨包调用
var Router router

// 声明router结构体
type router struct{}

func (r *router) InitApiRouter(router *gin.Engine) {
	router.
		// Query API
		GET("/get/questionlist", QuestionList.QueryQuestion).
		GET("/get/ranking", Ranking.QueryRanking).
		GET("/get/users", Users.QueryUsers).
		GET("/get/personal/ranking", Ranking.QueryPersonalRanking).
		// PUT FLag API
		PUT("/put/flag", Flag.PutFlag).
		// Create API
		POST("/create/question", QuestionList.CreateQuestion).
		//POST("/create/user", Users.CreateUsers).
		// Check API
		GET("/check/user", Users.CheckUsers)
}

// CORSMiddleware CORS中间件
// 防止跨域问题出现
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 设置允许的来源、HTTP方法和请求头
		c.Header("Access-Control-Allow-Origin", "*") // 替换为你的前端域名
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")

		// 处理预检请求
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
