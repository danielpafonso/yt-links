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

func (tmpl *Templates) ExecuteTemplate(data []Links) http.HandlerFunc {
	log.Println("hit execute template")
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("end execute template, %s\n", r.URL)
		err := tmpl.template.Execute(w, data)
		log.Println("end execute template")
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
		}
	}
}
