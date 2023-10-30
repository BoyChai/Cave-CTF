package controller

import (
	"Cave-CTF/dao"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

var QuestionList questionList

type questionList struct {
}

// CreateQuestion 插入数据
func (q *questionList) CreateQuestion(ctx *gin.Context) {
	var Question dao.QuestionList
	if err := ctx.Bind(Question); err != nil {
		fmt.Println("Bind请求参数失败, " + err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg":  err.Error(),
			"data": nil,
		})
		return
	}
	err := dao.Dao.CreateQuestion(Question)
	if err != nil {
		fmt.Println("插入数据失败, " + err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg":  err.Error(),
			"data": nil,
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"msg":  "插入数据成功",
		"data": nil,
	})
}

// DeleteQuestion 删除数据
//func (q *questionList) DeleteQuestion(ctx *gin.Context) {
//	var Question dao.QuestionList
//	if err := ctx.Bind(Question); err != nil {
//		fmt.Println("Bind请求参数失败, " + err.Error())
//		ctx.JSON(http.StatusInternalServerError, gin.H{
//			"msg":  err.Error(),
//			"data": nil,
//		})
//		return
//	}
//	err := dao.Dao.DeleteQuestion(int(Question.ID))
//	if err != nil {
//		fmt.Println("删除失败, " + err.Error())
//		ctx.JSON(http.StatusInternalServerError, gin.H{
//			"msg":  err.Error(),
//			"data": nil,
//		})
//		return
//	}
//	ctx.JSON(http.StatusOK, gin.H{
//		"msg":  "删除数据成功",
//		"data": nil,
//	})
//}

// QueryQuestion 查询数据
func (q *questionList) QueryQuestion(ctx *gin.Context) {
	err, lists := dao.Dao.FindQuestionList()
	if err != nil {
		fmt.Println("查询失败, " + err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg":  err.Error(),
			"data": nil,
		})
		return
	}
	for i := range lists {
		lists[i].Flag = ""
	}
	ctx.JSON(http.StatusOK, gin.H{
		"msg":  "查询成功",
		"data": lists,
	})
}
