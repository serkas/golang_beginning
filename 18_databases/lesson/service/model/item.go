package model

import "github.com/uptrace/bun"

// https://bun.uptrace.dev/guide/models.html#mapping-tables-to-structs

type Item struct {
	bun.BaseModel `bun:"table:items,alias:u" json:"-"`

	ID   int    `bun:"id,pk,autoincrement" json:"id"`
	Name string `bun:"name" json:"name"`
}
