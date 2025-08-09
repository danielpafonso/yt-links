package main

import (
	"flag"
	"log"
	"net/http"

	"youtube-links/internal"
)

func main() {
	// command line arguments
	var serverPort string
	var serverAddress string
	var templatePath string

	flag.StringVar(&serverAddress, "a", "0.0.0.0", "Server Address")
	flag.StringVar(&serverPort, "p", "8080", "Server Port")
	flag.StringVar(&templatePath, "t", "template/index.html.tmpl", "Template Path")
	flag.Parse()

	// "DB" handler
	data, err := internal.ReadStorage("storage.json")
	if err != nil {
		panic(err)
	}

	// Read template
	template, err := internal.LoadTemplate(templatePath)
	if err != nil {
		panic(err)
	}

	// Server handlers
	// webpage
	rootRouter := http.NewServeMux()
	rootRouter.HandleFunc("/", template.ExecuteTemplate(data))

	// api endpoints
	rootRouter.HandleFunc("POST /links", data.InsertData)
	rootRouter.HandleFunc("DELETE /links/{id}", data.DeleteById)

	server := http.Server{
		Addr:    serverAddress + ":" + serverPort,
		Handler: rootRouter,
	}

	log.Println("Server serving in port: " + serverAddress + ":" + serverPort)
	server.ListenAndServe()
}
