package models

type Todo struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	DueDate   string `json:"dueDate"`
	Completed bool   `json:"completed"`
}

type Todos struct {
	Todos []Todo `json:"todos"`
}
