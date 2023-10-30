package util

import (
	"encoding/json"
	"os"
)

func ReadEntityFile(file string, entity interface{}) error {
	bytes, err := os.ReadFile(file)
	if err != nil {
		return err
	}
	err = json.Unmarshal(bytes, entity)
	if err != nil {
		return err
	}
	return nil
}
