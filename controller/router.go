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
		// PUT FLag API
		PUT("/put/flag", Flag.PutFlag).
		// Create API
		POST("/create/question", QuestionList.CreateQuestion).
		POST("/create/user", Users.CreateUsers)
	// Delete API
}
