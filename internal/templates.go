package internal

import (
	"html/template"
	"log"
	"net/http"
)

type Templates struct {
	template *template.Template
}

func LoadTemplate(path string) (*Templates, error) {
	tmpt, err := template.ParseFiles(path)
	if err != nil {
		return nil, err
	}
	return &Templates{template: tmpt}, nil
}

func (tmpl *Templates) ExecuteTemplate(data mapLink) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Webpage request by: %s - %s\n", r.RemoteAddr, r.RequestURI)
		err := tmpl.template.Execute(w, data)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
		}
	}
}
