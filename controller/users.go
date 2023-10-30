package controller

import (
	"Cave-CTF/dao"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

var Users users

type users struct {
}

// CreateUsers 插入数据
func (q *users) CreateUsers(ctx *gin.Context) {
	var User dao.Users
	if err := ctx.Bind(&User); err != nil {
		fmt.Println("Bind请求参数失败, " + err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg":  err.Error(),
			"data": nil,
		})
		return
	}
	err := dao.Dao.CreateUsers(User)
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

// DeleteUsers 删除数据
//func (q *users) DeleteUsers(ctx *gin.Context) {
//	var User dao.Users
//	if err := ctx.Bind(User); err != nil {
//		fmt.Println("Bind请求参数失败, " + err.Error())
//		ctx.JSON(http.StatusInternalServerError, gin.H{
//			"msg":  err.Error(),
//			"data": nil,
//		})
//		return
//	}
//	err := dao.Dao.DeleteUsers(int(User.ID))
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

// QueryUsers 查询数据
func (q *users) QueryUsers(ctx *gin.Context) {
	err, lists := dao.Dao.FindUsers()
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

func (q *users) CheckUsers(ctx *gin.Context) {
	params := new(struct {
		Name  string `form:"name" binding:"required"`
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
	err, _ := dao.Dao.FindUser(params.Name, params.Alias)
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
		"data": true,
	})
}
