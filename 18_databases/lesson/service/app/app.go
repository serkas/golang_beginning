package app

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"proj/lessons/18_databases/lesson/service/model"
	"proj/lessons/18_databases/lesson/service/storage"
	"time"
)

type ItemsService interface {
	ListItems(ctx context.Context) ([]*model.Item, error)
	AddItem(ctx context.Context, item *model.Item) error
}

type App struct {
	conf   Config
	items  ItemsService
	server *http.Server
}

func New(conf Config) (*App, error) {
	return &App{
		conf: conf,
	}, nil
}

func (a *App) Run(ctx context.Context) error {
	sqlDB, err := sql.Open("mysql", a.conf.DB)
	if err != nil {
		return fmt.Errorf("db initialization: %w", err)
	}
	defer sqlDB.Close()

	app := &App{
		items: storage.New(sqlDB),
	}

	handler := http.NewServeMux()
	handler.HandleFunc("GET /items", app.getItems)
	handler.HandleFunc("POST /items", app.addItem)

	a.server = &http.Server{
		Handler: handler,
		Addr:    a.conf.ServerAddress,
	}

	go func() {
		<-ctx.Done()

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
		defer cancel()

		err := a.server.Shutdown(ctx)
		if err != nil {
			log.Printf("shutdonw: %s", err)
		}
	}()

	err = a.server.ListenAndServe()
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		return err
	}

	return nil
}

func (a *App) getItems(w http.ResponseWriter, r *http.Request) {
	items, err := a.items.ListItems(r.Context())
	if err != nil {
		log.Printf("getting items: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if items == nil {
		items = []*model.Item{}
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
