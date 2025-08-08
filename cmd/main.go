package main

import (
	"flag"
	"log"
	"net/http"

	"youtube-links/internal"
)

func serveWebpage(w http.ResponseWriter, r *http.Request) {
	log.Println("Webpage request by: " + r.RemoteAddr)
	w.Write([]byte("Hello, you hit our general endpoint"))
}

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
	dt, err := internal.ReadStorage("storage.json")
	if err != nil {
		panic(err)
	}
	log.Printf("%+v\n", dt)

	data := []internal.Links{
		{
			ID:   "g",
			Text: "google",
			Link: "https://www.google.com",
		},
		{
			ID:   "y",
			Text: "youtube",
			Link: "https://www.youtube.com",
		},
	}
	// err = internal.WriteStorage("storage.json", data)
	// if err != nil {
	// 	panic(err)
	// }

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
	// apiRouter := http.NewServeMux()
	// apiRouter.HandleFunc("POST /", apiHandler.InsertData)
	// apiRouter.HandleFunc("DELETE /{id}", apiHandler.DeleteById)
	// // add api sub routing
	// rootRouter.Handle("/links/", http.StripPrefix("/links", apiRouter))

	server := http.Server{
		Addr:    serverAddress + ":" + serverPort,
		Handler: rootRouter,
	}

	log.Println("Server serving in port: " + serverAddress + ":" + serverPort)
	server.ListenAndServe()
}
