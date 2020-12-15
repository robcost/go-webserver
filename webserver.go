package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	const index = "index.html"

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "" || r.URL.Path == "/" {
			http.ServeFile(w, r, index)
			log.Printf("Respond to request.")
		} else {
			http.ServeFile(w, r, r.URL.Path[1:])
			log.Printf("Respond to request.")
		}
	})

	log.Printf("Started webserver on port %s\n\n", port)
	http.ListenAndServe(":"+port, nil)
}
