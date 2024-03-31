package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

const serverPort = 8000

func main() {
	http.HandleFunc("/", rootHandler)

	http.HandleFunc("/about", aboutHandler)

	http.HandleFunc("/echo", echoHandler)

	log.Println("Starting server on port", serverPort)

	err := http.ListenAndServe(fmt.Sprintf(":%d", serverPort), nil)
	log.Fatal(err)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there! %s", r.URL.Path[1:])
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>About</h1> <p>It's a simple web page about this tiny web server</p>")
}

func echoHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("failed to read body: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	echoData := append([]byte("got: "), reqBody...)

	_, err = w.Write(echoData)
	if err != nil {
		log.Printf("failed to write response: %s", err)
	}
}
