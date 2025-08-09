package internal

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

type Link struct {
	Text string `json:"text"`
	Link string `json:"link"`
}

type postLink struct {
	Id   string `json:"id"`
	Text string `json:"text"`
	Link string `json:"link"`
}

type mapLink map[string]Link

// File Operations
func ReadStorage(path string) (mapLink, error) {
	data := make(map[string]Link)
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

func WriteStorage(path string, data mapLink) error {
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

// API Operations
// POST
func (mpl *mapLink) InsertData(w http.ResponseWriter, r *http.Request) {
	// read body
	body, err := io.ReadAll(r.Body)
	if err != nil || len(body) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(http.StatusText(http.StatusBadRequest)))
		return
	}

	fmt.Println(string(body))
}

// DELETE
func (mpl *mapLink) DeleteById(w http.ResponseWriter, r *http.Request) {
	// read path
	requestId := r.PathValue("id")
	fmt.Printf("delete id: %s\n", requestId)
}
