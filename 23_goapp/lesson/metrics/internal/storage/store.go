package storage

import (
	"context"
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/mysqldialect"

	"proj/lessons/23_goapp/lesson/metrics/internal/model"
)

type Store struct {
	db *bun.DB
}

func New(sqldb *sql.DB) *Store {
	return &Store{
		db: bun.NewDB(sqldb, mysqldialect.New()),
	}
}

func (s *Store) ListItems(ctx context.Context) (result []*model.Item, err error) {
	defer durationMetric("list_items")()

	err = s.db.NewSelect().Model(&result).Limit(1000).Scan(ctx)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *Store) GetItem(ctx context.Context, id int) (*model.Item, error) {
	defer durationMetric("get_item")()

	var item model.Item
	err := s.db.NewSelect().Model(&item).Where("id = ?", id).Scan(ctx)
	return &item, err
}

func (s *Store) AddItem(ctx context.Context, item *model.Item) error {
	defer durationMetric("add_item")()

	_, err := s.db.NewInsert().Model(item).Exec(ctx)

	return err
}

func (s *Store) GetTopLikedItems(ctx context.Context, limit int) (result []*model.Item, err error) {
	defer durationMetric("get_top_liked")()

	defer func(start time.Time) {
		log.Printf("got top liked from DB in %v", time.Since(start))
	}(time.Now())

	err = s.db.NewSelect().Model(&result).Order("likes DESC").Limit(limit).Scan(ctx)
	if err != nil {
		return nil, err
	}

	return result, nil
}
