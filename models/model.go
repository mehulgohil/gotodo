package models

type Todo struct {
	ID        int    `json:"id" ,yaml:"id"`
	Name      string `json:"name" ,yaml:"name"`
	DueDate   string `json:"dueDate" ,yaml:"due_date"`
	Completed bool   `json:"completed" ,yaml:"completed"`
}

type Todos struct {
	Todos []Todo `json:"todos" ,yaml:"todos"`
}
