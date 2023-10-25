package controller

import (
	"Cave-CTF/dao"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

var Flag flag

type flag struct {
}

func (f *flag) PutFlag(ctx *gin.Context) {
	params := new(struct {
		ID    int    `form:"question_id"`
		User  string `form:"user"`
		Score string `form:"score"`
		flag  string `form:"flag"`
	})
	if err := ctx.Bind(params); err != nil {
		fmt.Println("Bind请求参数失败, " + err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg":  err.Error(),
			"data": nil,
		})
		return
	}
	err, question := dao.Dao.FindQuestion(params.ID)
	if err != nil {
		fmt.Println("查询题目出现错误, " + err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg":  err.Error(),
			"data": nil,
		})
		return
	}
	if question.Flag == params.flag {
		err = dao.Dao.PutFlag(dao.Ranking{
			User:          params.User,
			Score:         params.Score,
			QuestionTitle: question.Title,
		})
		if err != nil {
			fmt.Println("增加分数出现错误, " + err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"msg":  err.Error(),
				"data": nil,
			})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"msg":  "提交成功",
			"data": nil,
		})
		return
	}
	ctx.JSON(http.StatusInternalServerError, gin.H{
		"msg":  "提交的flag错误" + err.Error(),
		"data": nil,
	})

}
