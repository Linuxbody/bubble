package models

import(
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// 链接 mysql
var DB *gorm.DB

func InitDB() (err error){
	// 链接 mysql
	DB, err = gorm.Open("mysql", "root:ts123456@tcp(127.0.0.1:3306)/bubble?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	} else {
		// 全局禁用表名复数
		DB.SingularTable(true) // 如果设置为true,`Todo`的默认表名为`Todo`,使用`TableName`设置的表名不受影响

		// 一般不会直接用CreateTable创建表
		// 检查模型`Todo`表是否存在，否则为模型`Todo`创建表
		if !DB.HasTable(&Todo{}) {
			if err := DB.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&Todo{}).Error; err !=nil {
				panic(err)
			}
		}
	}
	return	
}