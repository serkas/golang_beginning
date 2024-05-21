package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// Tutorial: https://tutorialedge.net/golang/golang-mysql-tutorial/

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:13306)/test_db")
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	results, err := db.Query("SHOW TABLES;")
	if err != nil {
		log.Fatal(err)
	}

	for results.Next() {
		var tableName string
		// for each row, scan the result into our tag composite object
		err = results.Scan(&tableName)
		if err != nil {
			log.Fatal(err)
		}
		// and then print out the tag's Name attribute
		fmt.Println("TABLE:", tableName)
	}

}
