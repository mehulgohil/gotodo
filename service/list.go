package service

import (
	"encoding/json"
	"errors"
	"github.com/mehulgohil/gotodo/models"
	"os"
	"strings"
)

func GetAllTodoTasks() (models.Todos, error) {
	currentTodosJson, err := os.ReadFile("storage.json")
	if err != nil && !strings.Contains(err.Error(), "The system cannot find the file specified") {
		return models.Todos{}, err
	}

	if len(currentTodosJson) < 1 {
		return models.Todos{}, errors.New("There are no todo tasks added.")
	}

	var currentTodos models.Todos
	err = json.Unmarshal(currentTodosJson, &currentTodos)
	if err != nil {
		return models.Todos{}, err
	}

	return currentTodos, nil
}
