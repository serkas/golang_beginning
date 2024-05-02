package storage

import (
	"context"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/mysqldialect"
	"log"
	"proj/lessons/19_redis/lesson/cached/internal/model"
	"time"
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
	err = s.db.NewSelect().Model(&result).Limit(1000).Scan(ctx)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *Store) GetItem(ctx context.Context, id int) (item *model.Item, err error) {
	err = s.db.NewSelect().Model(&item).Where("id = ?", id).Scan(ctx)
	return item, err
}

func (s *Store) AddItem(ctx context.Context, item *model.Item) error {
	_, err := s.db.NewInsert().Model(item).Exec(ctx)

	return err
}

func (s *Store) GetTopViewedItems(ctx context.Context, limit int) (result []*model.Item, err error) {
	defer func(start time.Time) {
		log.Printf("got top viewed from DB in %v", time.Since(start))
	}(time.Now())

	err = s.db.NewSelect().Model(&result).Order("views DESC").Limit(limit).Scan(ctx)
	if err != nil {
		return nil, err
	}

	return result, nil
}
