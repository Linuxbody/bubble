package models

// 定义 model
type Todo struct{
	ID int64 `json:"id"`
	Title string `json:"title"`
	Status bool `json:"status"`
}