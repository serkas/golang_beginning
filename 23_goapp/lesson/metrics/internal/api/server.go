package api

import (
	"context"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"

	"proj/lessons/23_goapp/lesson/metrics/internal/model"
)

type ItemsService interface {
	List(ctx context.Context) ([]*model.Item, error)
	Get(ctx context.Context, id int) (*model.Item, error)
	Add(ctx context.Context, item *model.Item) error
	GetTopLiked(ctx context.Context, limit int) ([]*model.Item, error)

	CountView(ctx context.Context, itemID int) error
	GetTopViewed(ctx context.Context, num int) (viewed []*model.Item, err error)
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
	handler.HandleFunc("GET /api/v1/items", metricsMiddleware(s.listItems))
	handler.HandleFunc("POST /api/v1/items", metricsMiddleware(s.addItem))
	handler.HandleFunc("GET /api/v1/items/{id}", metricsMiddleware(s.getItem))

	handler.HandleFunc("GET /api/v1/items/ranks/top_liked", metricsMiddleware(s.getTopLikedItems))
	handler.HandleFunc("GET /api/v1/items/ranks/top_viewed", metricsMiddleware(s.getTopViewedItems))

	// https://prometheus.io/docs/guides/go-application/
	handler.Handle("/metrics", promhttp.Handler())

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
