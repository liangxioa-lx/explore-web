package main

import (
	"explore/config"
	"explore/model"
	"explore/router"
	"github.com/gin-gonic/gin"
)

func main() {
	// 开启热加载
	config.HotLoad()
	// 初始化数据库
	model.Init()

	// 生成服务
	r := gin.Default()

	//路由注册
	router.Register(r)

	// 服务启动
	r.Run(config.Conf.Server.Address + ":" + config.Conf.Server.Port)
}
