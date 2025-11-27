package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Printf("Server started on port %s\n", port)

	http.Handle("/", http.FileServer(http.Dir("./frontend")))

	http.ListenAndServe(":"+port, nil)
}