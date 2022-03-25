package models

type Task struct {
	ID        string `json:"id" gorm:"primaryKey"`
	Name      string `json:"name"`
	Content   string `json:"content"`
	Completed bool   `json:"completed"`
}

type Info struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}
