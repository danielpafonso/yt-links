package internal

import (
	"encoding/json"
	"errors"
	"os"
)

type Links struct {
	ID   string `json:"id"`
	Text string `json:"text"`
	Link string `json:"link"`
}

func ReadStorage(path string) ([]Links, error) {
	data := make([]Links, 0)
	// read data
	fdata, err := os.ReadFile(path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			// empty configuration
			return data, nil
		}
		return nil, err
	}
	// unmarshal data
	err = json.Unmarshal(fdata, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func WriteStorage(path string, data []Links) error {

	jsonString, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}
	err = os.WriteFile(path, jsonString, 0666)
	if err != nil {
		return err
	}
	return nil
}
