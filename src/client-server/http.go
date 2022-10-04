package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func getPort() string {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}
	fmt.Printf("{\"Source\": \"backend-http\", \"Status\": {\"Port\": \"%v\"}}\n", port)
	return ":" + port
}

func greet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write([]byte(fmt.Sprintf("[ %s ] Hello from PDF-Rotator\n", time.Now())))
}

func getHomePage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write([]byte("http home page requested"))
}

func getMessageTransfer(w http.ResponseWriter, r *http.Request) {
	msg := r.FormValue("Message")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write([]byte(msg))
}

func main() {
	http.HandleFunc("/greet", greet)
	http.HandleFunc("/homepage", getHomePage)
	http.HandleFunc("/message", getMessageTransfer)
	http.ListenAndServe(getPort(), nil)
}
