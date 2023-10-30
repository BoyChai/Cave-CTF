package controller

import (
	"Cave-CTF/dao"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

var Flag flag

type flag struct {
}

var lastSubmitTime = make(map[string]time.Time)

func (f *flag) PutFlag(ctx *gin.Context) {

	params := new(struct {
		ID    int    `form:"question_id" json:"question_id"`
		User  string `form:"user"`
		Score string `form:"score"`
		Flag  string `form:"flag"`
	})
	if err := ctx.Bind(&params); err != nil {
		fmt.Println("Bind请求参数失败, " + err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg":  err.Error(),
			"data": nil,
		})
		return
	}

	user := params.User

	lastTime, ok := lastSubmitTime[user]
	if ok && time.Since(lastTime).Seconds() < 1 { // 限制 60 秒内只能提交一次
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg":  "请勿快速重复提交FLAG",
			"data": nil,
		})
		return
	}

	err, question := dao.Dao.FindQuestion(params.ID)
	if err != nil {
		lastSubmitTime[user] = time.Now()
		fmt.Println("查询题目出现错误, " + err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg":  err.Error(),
			"data": nil,
		})
		return
	}
	_, r := dao.Dao.FindRanking(params.User, question.Title)
	r2 := dao.Ranking{}
	fmt.Println(r)
	if r != r2 {

		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg":  "请勿重新提交FLAG",
			"data": nil,
		})
		return
	}
	if question.Flag == params.Flag {
		err = dao.Dao.PutFlag(dao.Ranking{
			User:          params.User,
			Score:         params.Score,
			QuestionTitle: question.Title,
		})
		if err != nil {
			lastSubmitTime[user] = time.Now()
			fmt.Println("增加分数出现错误, " + err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"msg":  err.Error(),
				"data": nil,
			})
			return
		}
		lastSubmitTime[user] = time.Now()
		ctx.JSON(http.StatusOK, gin.H{
			"msg":  "提交成功",
			"data": nil,
		})
		return
	}
	ctx.JSON(http.StatusInternalServerError, gin.H{
		"msg":  fmt.Sprintf("错误的flag"),
		"data": nil,
	})
	lastSubmitTime[user] = time.Now()

}
