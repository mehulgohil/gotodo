package service

import (
	"encoding/json"
	"errors"
	"github.com/mehulgohil/gotodo/models"
	"os"
	"strings"
)

func SearchTodoTask(nameKeyword string, duedate string) (models.Todos, error) {
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

	if nameKeyword == "" && duedate == "" {
		return currentTodos, nil
	}

	var finalList models.Todos
	for _, eachTask := range currentTodos.Todos {
		if duedate != "" && eachTask.DueDate == duedate && strings.Contains(eachTask.Name, nameKeyword) {
			finalList.Todos = append(finalList.Todos, eachTask)
		} else if duedate == "" && strings.Contains(eachTask.Name, nameKeyword) {
			finalList.Todos = append(finalList.Todos, eachTask)
		}
	}

	return finalList, nil
}
