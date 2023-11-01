package util

import (
	"bufio"
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

func Scan(path string, callback func(line string)) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		callback(scanner.Text())
	}
	return nil
}
