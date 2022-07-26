package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	// Get port from environment
	port, ok := os.LookupEnv("PORT")
	if !ok {
		port = "8080"
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello world")
	})


	// Run the server
	// This will block the current thread
	http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
}
