package api

import (
	"context"
	"log"
	"net"
	"net/http"
	"proj/lessons/21_di/lesson/service/internal/model"
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
	handler.HandleFunc("GET /items", s.listItems)
	handler.HandleFunc("POST /items", s.addItem)
	handler.HandleFunc("GET /items/{id}", s.getItem)

	handler.HandleFunc("GET /items/ranks/top_liked", s.getTopLikedItems)
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

func (s *Server) Start() error {
	ln, err := net.Listen("tcp", s.server.Addr)
	if err != nil {
		return err
	}

	log.Println("Starting HTTP server at", s.server.Addr)
	go s.server.Serve(ln)

	return nil
}
