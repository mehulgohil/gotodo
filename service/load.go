package service

import (
	"encoding/json"
	"github.com/mehulgohil/gotodo/models"
	"os"
)

func LoadExistingJSON(todos models.Todos) error {
	marshalJson, err := json.Marshal(todos)
	if err != nil {
		return err
	}

	err = os.WriteFile("storage.json", marshalJson, 0644)
	if err != nil {
		return err
	}

	return nil
}
