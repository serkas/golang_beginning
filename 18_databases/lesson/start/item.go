package main

type Item struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
}
