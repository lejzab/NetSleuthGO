package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s\n", time.Now().Format("2006-01-02 15:04:05"), r.Method, r.URL.Path)
		fmt.Fprintf(w, "Hello, world!")
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
	log.Println("Server started on http://localhost:8080")

}
