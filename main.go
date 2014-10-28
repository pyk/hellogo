package main

import (
	"fmt"
	"log"
	"net/http"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}
func main() {
	http.HandleFunc("/", HelloHandler)

	log.Println("start listening on port 8080..")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
