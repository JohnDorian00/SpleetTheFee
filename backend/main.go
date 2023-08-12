package main

import (
	"fmt"
	"log"
	"net/http"

	exp "backend/testing"
)

func main() {
	// Запуск кода из полигона
	exp.Main()
	//_____________________________________________________________________
	http.HandleFunc("/hello", HelloHandler)
	log.Println("Listening...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func HelloHandler(w http.ResponseWriter, _ *http.Request) {

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Max-Age", "15")
	fmt.Fprintf(w, "Hello, there!")
}
