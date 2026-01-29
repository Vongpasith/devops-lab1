package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello DevOps Lab 1")
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // default
	}

	http.HandleFunc("/", handler)

	log.Println("Server runing on :" + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
