package main

import (
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("./"))

	http.Handle("/", http.StripPrefix("/", fs))

	log.Println("drawdb server running on :3000")
	if err := http.ListenAndServe(":3000", nil); err != nil {
		log.Fatalf("could not start server: %s\n", err)
	}
}
