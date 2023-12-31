package service

import (
	"encoding/json"
	"github.com/mehulgohil/gotodo/models"
	"os"
	"strings"
)

func AddTodoTask(todo models.Todo) error {
	currentTodosJson, err := os.ReadFile("storage.json")
	if err != nil && !strings.Contains(err.Error(), "The system cannot find the file specified") {
		return err
	}

	todo.ID = 1
	var currentTodos models.Todos
	if len(currentTodosJson) > 1 {
		err = json.Unmarshal(currentTodosJson, &currentTodos)
		if err != nil {
			return err
		}

		todo.ID = len(currentTodos.Todos) + 1
	}

	currentTodos.Todos = append(currentTodos.Todos, todo)

	marshalJson, err := json.Marshal(currentTodos)
	if err != nil {
		return err
	}

	err = os.WriteFile("storage.json", marshalJson, 0644)
	if err != nil {
		return err
	}

	return nil
}
