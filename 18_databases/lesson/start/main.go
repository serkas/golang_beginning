package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	"github.com/jmoiron/sqlx"
	"log"
)

func main() {
	//demoStandardLib()
	//os.Exit(0)

	// Open a MySQL database.
	sqldb, err := sqlx.Open("mysql", "root:root@tcp(localhost:13306)/start")
	if err != nil {
		log.Fatal(err)
	}

	//sqldb.MustExec(`
	//	CREATE TABLE items2 (
	//	id   BIGINT NOT NULL AUTO_INCREMENT,
	//	name varchar(250) NOT NULL,
	//	PRIMARY KEY (id));
	//`)

	// Insert
	newItem := Item{
		Name: "New Item 4",
	}

	_, err = sqldb.NamedExec("INSERT INTO items (name) VALUES (:name)", &newItem)
	if err != nil {
		log.Fatal(err)
	}

	// Select one
	items := []*Item{}
	err = sqldb.Select(&items, "SELECT * FROM items WHERE id = ? ORDER BY id ASC", 11)
	if err != nil {
		log.Fatal(err)
	}

	for _, item := range items {
		fmt.Println(item)
	}

	//tx, err := sqldb.Begin()
	//
	//_, err = tx.Exec(``)
	//if err != nil {
	//
	//}
	//
	//err = tx.Commit()

}

func demoStandardLib() {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:13306)/start")
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	results, err := db.Query("SELECT id, name FROM items;")
	if err != nil {
		log.Fatal(err)
	}

	for results.Next() {
		var item Item
		// for each row, scan the result into our tag composite object
		err = results.Scan(&item.ID, &item.Name)
		if err != nil {
			log.Fatal(err)
		}
		// and then print out the tag's Name attribute
		fmt.Println(item)
	}
}
