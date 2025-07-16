package main

import (
	"clouddisk/lib"
	"clouddisk/model/mysql"
	"clouddisk/router"
	"fmt"
	"log"
)

func main() {
	fmt.Println("启动云盘")
	//	加载服务配置
	serverConfig := lib.LoadServerConfig()

	mysql.InitDB(serverConfig)

	r := router.SetupRoute()
	//注册HTML模板，从view/目录加载所有模板文件（用于服务端渲染）
	r.LoadHTMLGlob("view/*")
	//注册HTML模板，从view/目录加载所有模板文件（用于服务端渲染）
	r.Static("/static", "./static")

	if err := r.Run(":80"); err != nil {
		log.Fatal("服务器启动失败...")
	}
	defer mysql.DB.Close()
	fmt.Println("hh")
	//初始化数据库
	//设置路由
	//加载网页资源
	//启动服务器
}
