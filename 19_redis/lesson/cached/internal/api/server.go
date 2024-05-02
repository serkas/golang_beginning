package api

import (
	"context"
	"net/http"
	"proj/lessons/19_redis/lesson/cached/internal/model"
)

type ItemsService interface {
	List(ctx context.Context) ([]*model.Item, error)
	Get(ctx context.Context, id int) (*model.Item, error)
	Add(ctx context.Context, item *model.Item) error
	GetTopViewed(ctx context.Context, limit int) ([]*model.Item, error)
}

type Server struct {
	items  ItemsService
	server *http.Server
}

func NewServer(addr string, items ItemsService) *Server {
	s := &Server{
		items: items,
	}

	handler := http.NewServeMux()
	handler.HandleFunc("GET /items", s.getItems)
	handler.HandleFunc("POST /items", s.addItem)

	handler.HandleFunc("GET /items/ranks/top_viewed", s.getTopViewedItems)

	s.server = &http.Server{
		Handler: handler,
		Addr:    addr,
	}

	return s
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}

func (s *Server) Serve() error {
	return s.server.ListenAndServe()
}
