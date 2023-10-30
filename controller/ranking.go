package controller

import (
	"Cave-CTF/dao"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

var Ranking ranking

type ranking struct {
}

// QueryRanking 查询数据
func (r *ranking) QueryRanking(ctx *gin.Context) {
	err, lists := dao.Dao.FindRankings()
	if err != nil {
		fmt.Println("查询失败, " + err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg":  err.Error(),
			"data": nil,
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"msg":  "查询成功",
		"data": lists,
	})
}

func (r *ranking) QueryPersonalRanking(ctx *gin.Context) {
	params := new(struct {
		Alias string `form:"alias" binding:"required"`
	})
	if err := ctx.Bind(&params); err != nil {
		fmt.Println("Bind请求参数失败, " + err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg":  err.Error(),
			"data": nil,
		})
		return
	}
	fmt.Println(params.Alias)
	_, rankings := dao.Dao.FindPersonalRanking(params.Alias)
	//if err != nil {
	//	fmt.Println("查询失败, " + err.Error())
	//	ctx.JSON(http.StatusInternalServerError, gin.H{
	//		"msg":  err.Error(),
	//		"data": nil,
	//	})
	//	return
	//}
	ctx.JSON(http.StatusOK, gin.H{
		"msg":  "查询成功",
		"data": rankings,
	})
}
