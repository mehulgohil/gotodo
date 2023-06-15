package service

import (
	"encoding/json"
	"errors"
	"github.com/mehulgohil/gotodo/models"
	"os"
)

func RemoveTodoTask(taskId int) error {
	if taskId <= 0 {
		return errors.New("invalid task id")
	}

	currentTodosJson, err := os.ReadFile("storage.json")
	if err != nil {
		return err
	}

	if len(currentTodosJson) < 1 {
		return errors.New("There are no todo tasks added.")
	}

	var currentTodos models.Todos
	err = json.Unmarshal(currentTodosJson, &currentTodos)
	if err != nil {
		return err
	}

	if len(currentTodos.Todos) < taskId {
		return errors.New("Todo Task with given Task ID doesn't exist")
	}

	currentTodos.Todos = append(currentTodos.Todos[:taskId-1], currentTodos.Todos[taskId:]...)

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
