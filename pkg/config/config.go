package config

import (
	"os"
	"errors"
	"encoding/json"
)

func Load(path string, model interface{}) (error) {
	if path == "" {
		return errors.New("Missing path")
	}

	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(model)
	if err != nil {
		return err
	}

	return nil
}