package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

var (
	Version = ""
)

func main() {
	addr, ok := os.LookupEnv("SERVER_ADDR")
	if !ok {
		addr = ":8081"
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(fmt.Sprintf("%v", time.Now())))
	})

	http.HandleFunc("/info", func(w http.ResponseWriter, r *http.Request) {
		info := struct {
			Version string
		}{
			Version: Version,
		}

		err := json.NewEncoder(w).Encode(info)
		if err != nil {
			log.Println(err)
		}
	})

	log.Printf("starting server at %s", addr)

	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatal(err)
	}
}
