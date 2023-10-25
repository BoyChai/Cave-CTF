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
	err, lists := dao.Dao.FindRanking()
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
