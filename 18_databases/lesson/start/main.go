package main

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/mysqldialect"
	"log"
)

func main() {
	//demoStandardLib()
	//os.Exit(0)

	// Open a MySQL database.
	sqldb, err := sql.Open("mysql", "root:root@tcp(localhost:13306)/start")
	if err != nil {
		log.Fatal(err)
	}

	// Create a Bun db on top of it.
	db := bun.NewDB(sqldb, mysqldialect.New())

	// Select one
	var item Item
	err = db.NewSelect().Model(&item).Where("id = ?", 2).Scan(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(item)

	// Insert
	newItem := Item{
		Name: "New Item",
	}
	_, err = db.NewInsert().Model(&newItem).Exec(context.Background())
	if err != nil {
		log.Fatal(err)
	}

}

func demoStandardLib() {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:13306)/start")
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	results, err := db.Query("SELECT * FROM items;")
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
