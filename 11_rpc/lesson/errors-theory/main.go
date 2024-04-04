package main

import (
	"errors"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"proj/lessons/11_rpc/lesson/errors-theory/service"
)

func main() {
	addr := ":8003"
	slog.Info("starting server", "addr", addr)
	if err := http.ListenAndServe(addr, &handler{}); err != nil {
		log.Fatal(err)
	}
}

type handler struct {
}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	caseParameter := r.URL.Query().Get("case")

	err := someLayerOfExecution(caseParameter)
	if err != nil {
		handleErr(w, err)
		return
	}
}

func someLayerOfExecution(param string) error {
	err := service.Process(param)
	if err != nil {
		return fmt.Errorf("processing: %w", err)
	}

	return nil
}

func handleErr(w http.ResponseWriter, err error) {
	if errors.Is(err, service.ErrAlreadyExist) {
		slog.Warn("already exist error", "err", err)
		w.WriteHeader(http.StatusConflict)
		return
	}

	var notFoundErr *service.NotFoundError
	if errors.As(err, &notFoundErr) {
		slog.Warn("entity not found", "err", err, "id", notFoundErr.ID)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	slog.Warn("execution error", "err", err)
	w.WriteHeader(http.StatusInternalServerError)
}
