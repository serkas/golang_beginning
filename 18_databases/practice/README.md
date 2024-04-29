## Using sqlx

sqlx is a library which provides a set of extensions on go's standard database/sql library

https://github.com/jmoiron/sqlx


go get  "github.com/jmoiron/sqlx"

### simple table schema

```mysql
CREATE TABLE items (
    id   BIGINT NOT NULL AUTO_INCREMENT,
    name varchar(250) NOT NULL,
    PRIMARY KEY (id)
);
```

Create it with your SQL client or directly in go program (it's ok for this practice):

```go
	sqldb, err := sqlx.Open("mysql", "root:root@tcp(localhost:13306)/your_db_name")
	if err != nil {
		log.Fatal(err)
	}

	sqldb.MustExec(`
		CREATE TABLE items (
		id   BIGINT NOT NULL AUTO_INCREMENT,
		name varchar(250) NOT NULL,
		PRIMARY KEY (id));
	`)
```

### insert

```go
	newItem := Item{Name: "New Item 4"}
	_, err = sqldb.NamedExec("INSERT INTO items (name) VALUES (:name)", &newItem)
	if err != nil {
		log.Fatal(err)
	}
```

### select

```go
	items := []*Item{}
	err = sqldb.Select(&items, "SELECT * FROM items ORDER BY id ASC")
	if err != nil {
		log.Fatal(err)
	}
```


