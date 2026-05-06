package storage

import (
	"encoding/json"
	"os"
)

func Save() error {
	file, err := os.Create("storage/data.json")
	if err != nil {
		return err
	}
	defer file.Close()
	return json.NewEncoder(file).Encode(Store)
}
func Load() error {
	file, err := os.Open("storage/data.json")
	if err != nil {
		return err
	}
	defer file.Close()
	return json.NewDecoder(file).Decode(&Store)
}
