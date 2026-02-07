package main 

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

/* Handlers */

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to the Shapes API"))
}

func health(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Server is running"))
}

func about(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Allen"))
}

func timeHandler(w http.ResponseWriter, r *http.Request) {
	now := time.Now().Format("2006-01-02 15:04:05")
	w.Write([]byte(now))
}

// Custom Route
func random(w http.ResponseWriter, r *http.Request) {
	num := rand.Intn(100)
	w.Write([]byte(fmt.Sprintf("%d", num)))
}

func main() {

	http.HandleFunc("/", home)
	http.HandleFunc("/health", health)
	http.HandleFunc("/about", about)
	http.HandleFunc("/time", timeHandler)
	http.HandleFunc("/random", random)

	fmt.Println("Server running on :4000")

	http.ListenAndServe(":4000", nil)
}
