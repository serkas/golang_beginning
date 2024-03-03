package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type book struct {
	Title      string `json:"title"`
	Author     string `json:"author"`
	Year       int    `json:"year"`
	TotalPages int    `json:"total_pages"`

	OrdersCount int `json:"-"`
}

var staticBook = book{
	Title:       "Dune",
	Author:      "Frank Herbert",
	Year:        1965,
	TotalPages:  896,
	OrdersCount: 123,
}

func main() {
	data, err := json.Marshal(staticBook)
	if err != nil {
		log.Fatalf("marshaling: %s", err)
	}

	http.HandleFunc("/book", func(w http.ResponseWriter, r *http.Request) {
		w.Write(data)
	})
	addr := ":8080"
	log.Printf("starting server at %s", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
