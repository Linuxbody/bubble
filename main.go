package main

import(
	"bubble/models"
	"bubble/routers"
)

// 前端启动
// npm install cnpm --registry=https://registry.npm.taobao.org
// npm run serve

func main(){
	// 创建数据库
	// CREATE DATABASE bubble;
	// 链接数据库
	err := models.InitDB()
	if err != nil{
		panic(err)
	}

	r := routers.SetupRouter()
	r.Run(":8000")
}
