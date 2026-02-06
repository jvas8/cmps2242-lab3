package main

import (
	"net/http"
	"time"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to the Shapes API\n"))
}

func health(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Server is running\n"))
}

func about(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Jeimy Vasquez\n"))
}

func timeHandler(w http.ResponseWriter, r *http.Request) {
	currentTime := time.Now().Format("2006-01-02 15:04:05\n")
	w.Write([]byte(currentTime))
}

func greeting(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello! Thanks for visiting the API\n"))
}

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/health", health)
	http.HandleFunc("/about", about)
	http.HandleFunc("/time", timeHandler)
	http.HandleFunc("/greeting", greeting)

	http.ListenAndServe(":4000", nil)
}
