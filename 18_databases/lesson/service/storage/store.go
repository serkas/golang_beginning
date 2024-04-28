package storage

import (
	"context"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/mysqldialect"
	"proj/lessons/18_databases/lesson/service/model"
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
	err = s.db.NewSelect().Model(&result).Scan(ctx)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *Store) AddItem(ctx context.Context, item *model.Item) error {
	_, err := s.db.NewInsert().Model(item).Exec(ctx)

	return err
}
