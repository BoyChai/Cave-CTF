package main

import (
	"Cave-CTF/config"
	"Cave-CTF/controller"
	"Cave-CTF/dao"
	"github.com/gin-gonic/gin"
)

func main() {

	config.Config.Read()
	dao.Client()
	dao.AutoTables()
	r := gin.Default()
	controller.Router.InitApiRouter(r)
	r.Run()
}
