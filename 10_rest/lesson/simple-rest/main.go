package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"simple-rest/entities"
	"simple-rest/store"
)

type SensorStorage interface {
	List() ([]*entities.Sensor, error)
	Get(id int64) (*entities.Sensor, error)
	Create(sensor *entities.Sensor) error
	Update(sensor *entities.Sensor) error
	Delete(id int64) error
}

func main() {
	srv := &Server{storage: store.NewMemStorage()}

	r := mux.NewRouter()

	r.HandleFunc("/api/sensors", srv.listHandler).Methods(http.MethodGet)
	r.HandleFunc("/api/sensors/{id}", srv.getHandler).Methods(http.MethodGet)
	r.HandleFunc("/api/sensors", srv.createHandler).Methods(http.MethodPost)
	r.HandleFunc("/api/sensors/{id}", srv.updateHandler).Methods(http.MethodPut)
	r.HandleFunc("/api/sensors/{id}", srv.deleteHandler).Methods(http.MethodDelete)

	if err := http.ListenAndServe(":8001", r); err != nil {
		log.Printf("server run error: %s", err)
	}
}

type Server struct {
	storage SensorStorage
}

func (s *Server) listHandler(w http.ResponseWriter, r *http.Request) {
	sensors, err := s.storage.List()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeResponse(w, sensors)
}

func (s *Server) getHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.ParseInt(params["id"], 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	sensor, err := s.storage.Get(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeResponse(w, sensor)
}

func (s *Server) createHandler(w http.ResponseWriter, r *http.Request) {
	var sensor entities.Sensor
	err := json.NewDecoder(r.Body).Decode(&sensor)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = s.storage.Create(&sensor)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (s *Server) updateHandler(w http.ResponseWriter, r *http.Request) {
	var sensor entities.Sensor
	err := json.NewDecoder(r.Body).Decode(&sensor)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	params := mux.Vars(r)
	id, err := strconv.ParseInt(params["id"], 10, 64)
	if err != nil || id != sensor.ID {
		http.Error(w, "invalid_id", http.StatusBadRequest)
		return
	}

	err = s.storage.Update(&sensor)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (s *Server) deleteHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.ParseInt(params["id"], 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = s.storage.Delete(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func writeResponse(w http.ResponseWriter, data any) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		log.Printf("writing response error: %s", err)
	}
}
