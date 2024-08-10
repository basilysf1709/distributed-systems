package store

import (
	"encoding/json"
	"os"
)

type JSONPersistence struct {
	FilePath string
}

func (jp *JSONPersistence) Save(data map[string]string) error {
	file, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(jp.FilePath, file, 0644)
}

func (jp *JSONPersistence) Load() (map[string]string, error) {
	file, err := os.ReadFile(jp.FilePath)
	if err != nil {
		if os.IsNotExist(err) {
			return make(map[string]string), nil
		}
		return nil, err
	}

	var data map[string]string
	err = json.Unmarshal(file, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}