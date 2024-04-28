package main

import "github.com/uptrace/bun"

// https://bun.uptrace.dev/guide/models.html#mapping-tables-to-structs

type Item struct {
	bun.BaseModel `bun:"table:items,alias:u"`

	ID   int    `bun:"id,pk,autoincrement"`
	Name string `bun:"name"`
}
