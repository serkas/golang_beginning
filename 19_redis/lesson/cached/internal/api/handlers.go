package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"proj/lessons/19_redis/lesson/cached/internal/model"
)

func (s *Server) getItems(w http.ResponseWriter, r *http.Request) {
	items, err := s.items.List(r.Context())
	if err != nil {
		log.Printf("getting items: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if items == nil {
		items = []*model.Item{}
	}

	writeResponse(w, items)
}

func (s *Server) addItem(w http.ResponseWriter, r *http.Request) {
	var item model.Item

	err := json.NewDecoder(r.Body).Decode(&item)
	if err != nil {
		log.Printf("unmarshaling item: %s", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = s.items.Add(r.Context(), &item)
	if err != nil {
		log.Printf("adding item: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (s *Server) getTopViewedItems(w http.ResponseWriter, r *http.Request) {
	numTopItems := 10
	items, err := s.items.GetTopViewed(r.Context(), numTopItems)
	if err != nil {
		handleError(w, fmt.Errorf("getting top viewed items: %w", err))
		return
	}

	if items == nil {
		items = []*model.Item{}
	}

	writeResponse(w, items)
}

func writeResponse(w http.ResponseWriter, data any) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		log.Printf("writing response: %s", err)
	}
}

func handleError(w http.ResponseWriter, err error) {
	log.Print(err)
	w.WriteHeader(http.StatusInternalServerError)
}
