package internal

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type Link struct {
	Text string `json:"text"`
	Link string `json:"link"`
}

type postLink struct {
	Id   string `json:"id"`
	Link string `json:"link"`
}

type mapLink map[string]Link

const (
	linkFullTemplate   string = "https://www.youtube.com/embed/%s?vq=hd720&start=%s"
	linkSimpleTemplate string = "https://www.youtube.com/embed/%s?vq=hd720"
)

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
func (mpl mapLink) InsertData(w http.ResponseWriter, r *http.Request) {
	log.Printf("POST request by: %s - %s\n", r.RemoteAddr, r.RequestURI)
	// read body
	body, err := io.ReadAll(r.Body)
	if err != nil || len(body) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(http.StatusText(http.StatusBadRequest)))
		return
	}

	// unmarshal body
	var linkBody Link
	err = json.Unmarshal(body, &linkBody)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(http.StatusText(http.StatusBadRequest)))
		return
	}

	// check if fields are empty
	if linkBody.Link == "" || linkBody.Text == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(http.StatusText(http.StatusBadRequest)))
		return
	}

	// parse link
	id, t, err := LinkParser(linkBody.Link)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(http.StatusText(http.StatusBadRequest)))
		return
	}
	// update link inplace
	if t == "" {
		linkBody.Link = fmt.Sprintf(linkSimpleTemplate, id)
	} else {
		linkBody.Link = fmt.Sprintf(linkFullTemplate, id, t)
	}

	// update storage
	mpl[id] = linkBody

	// send response
	rsp := postLink{
		Id:   id,
		Link: linkBody.Link,
	}
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(rsp)
}

// DELETE
func (mpl mapLink) DeleteById(w http.ResponseWriter, r *http.Request) {
	log.Printf("DELETE request by: %s - %s\n", r.RemoteAddr, r.RequestURI)
	// read path
	requestId := r.PathValue("id")
	if _, ok := mpl[requestId]; ok {
		delete(mpl, requestId)
		w.WriteHeader(http.StatusNoContent)
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}
