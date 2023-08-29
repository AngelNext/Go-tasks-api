package models

type Task struct {
	ID            string `json:"id" gorm:"primaryKey"`
	Title         string `json:"title"`
	Description   string `json:"description"`
  CreatedAt     int    `json:"created_at"`
	Done          bool   `json:"done"`
}
