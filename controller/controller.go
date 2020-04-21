package controller

import(
	"github.com/gin-gonic/gin"
	"net/http"
	"bubble/models"
)
func IndexHandler(c *gin.Context){
	c.JSON(http.StatusOK, gin.H{
		"Msg": "StatusOk",
	})
}

var err error
// 添加
func CreateHandler(c *gin.Context){
	// 前端页面填写待办事项， 点击提交 会发送请求到这里
	// 1、从请求中把数据拿出来
	var todo models.Todo
	c.BindJSON(&todo)

	// 2、存入数据库
	// err = DB.Create(&todo).Error
	// if err != nil{	
	// }
	// 3、返回相应
	// 存数据和返回响应一起判断
	if err = models.DB.Create(&todo).Error;err != nil{
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
	}else{
		c.JSON(http.StatusOK, todo)
	}
}

// 查看所有代办事项
func GetListHandler(c *gin.Context){
	var todo []models.Todo 
	models.DB.Find(&todo)
	c.JSON(http.StatusOK, &todo)

}

// 查看某一个代办事项
func GetTodoHandler(c *gin.Context){
	var todo models.Todo
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{"error": "无效的ID！"})
		return
	}
	if err = models.DB.First(&todo, id).Error; err != nil{
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	}else{
		c.JSON(http.StatusOK, &todo)
	}
}

// 修改
func UpdateTodoHandler(c *gin.Context){
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{"error": "无效的ID"})
		return
	}
	// 查询获取数据
	var todo models.Todo
	if err = models.DB.Where("id=?", id).First(&todo).Error; err !=nil{
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}
	// 绑定数据
	c.BindJSON(&todo)
	if err = models.DB.Save(todo).Error; err != nil{
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	}else{
		c.JSON(http.StatusOK, todo)
	}
}

// 删除某一个代办事项
func DeleteTodoHandler(c *gin.Context){
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{"error": "删除失败！"})
		return
	}
	if err = models.DB.Where("id=?", id).Delete(models.Todo{}).Error; err != nil{
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	}else{
		c.JSON(http.StatusOK, gin.H{id: "deleted"})
	}

}