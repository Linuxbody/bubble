package routers

import(
	"github.com/gin-gonic/gin"
	"bubble/controller"
)
func SetupRouter() *gin.Engine{
	r := gin.Default()
	// 告诉gin 框架模板文件引用的静态文件目录
	// r.Static("/static", "static")
	// // 告诉gin 框架模板文件目录
	// r.LoadHTMLGlob("/templates/*")

	r.GET("/", controller.IndexHandler)

	// 定义 v1 Api接口
	bubbleV1 := r.Group("/v1")
	{
		// 待办事项
		// 添加
		bubbleV1.POST("/todo", controller.CreateHandler)

		// 查看所有代办事项
		bubbleV1.GET("/todo", controller.GetListHandler)

		// 查看某一个代办事项
		bubbleV1.GET("/todo/:id", controller.GetTodoHandler)

		// 修改
		bubbleV1.PUT("/todo/:id", controller.UpdateTodoHandler)

		// 删除某一个代办事项
		bubbleV1.DELETE("/todo/:id", controller.DeleteTodoHandler)
	}
	return r
}