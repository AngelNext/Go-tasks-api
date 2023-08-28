package models

type Task struct {
	ID        string `json:"id" gorm:"primaryKey"`
	Title      string `json:"title"`
	Description   string `json:"description"`
	Done bool   `json:"done"`
}
