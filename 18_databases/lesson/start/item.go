package main

type Item struct {
	ID   int    `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
}
