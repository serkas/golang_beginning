package app

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"proj/lessons/16_testing/lesson/functional/model"
	"time"
)

type ItemsService interface {
	ListItems(ctx context.Context) ([]*model.Item, error)
	AddItem(ctx context.Context, item *model.Item) error
}

type App struct {
	items  ItemsService
	server *http.Server
}

func New(items ItemsService, serverAddress string) *App {
	app := &App{
		items: items,
	}

	handler := http.NewServeMux()
	handler.HandleFunc("GET /items", app.getItems)
	handler.HandleFunc("POST /items", app.addItem)

	app.server = &http.Server{
		Addr:    serverAddress,
		Handler: handler,
	}

	return app
}

func (a *App) Run(_ context.Context) error {
	err := a.server.ListenAndServe()
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		return err
	}

	return nil
}

func (a *App) Shutdown() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	err := a.server.Shutdown(ctx)
	if err != nil {
		log.Printf("shutdonw: %s", err)
	}
}

func (a *App) getItems(w http.ResponseWriter, r *http.Request) {
	items, err := a.items.ListItems(r.Context())
	if err != nil {
		log.Printf("getting items: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(items)
	if err != nil {
		log.Printf("writing response: %s", err)
	}
}

func (a *App) addItem(w http.ResponseWriter, r *http.Request) {
	var item model.Item

	err := json.NewDecoder(r.Body).Decode(&item)
	if err != nil {
		log.Printf("unmarshaling item: %s", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = a.items.AddItem(r.Context(), &item)
	if err != nil {
		log.Printf("adding item: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
