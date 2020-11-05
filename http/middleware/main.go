package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", myMiddleware(MyHandler))

	http.ListenAndServe(":8080", nil)
}

func MyHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("handler work")
	w.WriteHeader(http.StatusOK)
}

func myMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("Middleware work")
		next(w, r)
	}
}
