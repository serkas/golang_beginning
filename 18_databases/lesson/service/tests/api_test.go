package tests_test

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io"
	"net/http"
	"proj/lessons/18_databases/lesson/service/app"
	"proj/lessons/18_databases/lesson/service/model"
	"proj/lessons/18_databases/lesson/service/storage"
	"testing"
	"time"
)

func setup(t *testing.T) (cleanUp func(...string), db *sql.DB) {
	conf := app.Config{
		ServerAddress: ":8888",
		DB:            "root:root@tcp(localhost:23306)/items_db_test",
	}

	a, err := app.New(conf)
	if err != nil {
		t.Logf("app setup: %s", err)
	}

	ctx, cancelSrv := context.WithCancel(context.Background())
	go func() {
		err := a.Run(ctx)
		if err != nil {
			t.Logf("running server: %s", err)
		}
	}()

	// waiting until the server is up
	assert.Eventually(t, func() bool {
		_, err := http.Get("http://localhost:8888")
		return err == nil
	}, time.Second, 10*time.Millisecond)

	// DB client:
	// * to prepare fixtures
	// * to assert state
	// * to clean up after
	sqlDB, err := sql.Open("mysql", conf.DB)
	if err != nil {
		defer cancelSrv()
		t.Fatalf("test db connection: %s", err)
	}

	cleanUp = func(dbTables ...string) {
		cancelSrv()

		for _, table := range dbTables {
			_, err := sqlDB.Exec(fmt.Sprintf("TRUNCATE TABLE %s", table)) // DO NOT USE string formatting in app code. TRUNCATE does not support SQL arguments, so we use string operations
			if err != nil {
				t.Errorf("cleaning table %s: %s", table, err)
			}
		}

		sqlDB.Close()
	}

	return cleanUp, sqlDB
}

func TestAddItem(t *testing.T) {
	cleanUp, db := setup(t)
	defer cleanUp("items")

	// test-case setup
	item := &model.Item{
		ID:   2,
		Name: "item 2",
	}
	data, err := json.Marshal(item)
	require.NoError(t, err)

	req, err := http.NewRequest("POST", "http://localhost:8888/items", bytes.NewReader(data))
	require.NoError(t, err)

	// execution
	cli := http.DefaultClient
	resp, err := cli.Do(req)
	require.NoError(t, err)

	// validation
	require.Equal(t, http.StatusCreated, resp.StatusCode)

	store := storage.New(db)
	storedItems, err := store.ListItems(context.Background())
	require.NoError(t, err)
	require.Len(t, storedItems, 1)
	assert.Equal(t, item.Name, storedItems[0].Name)
}

func TestListItems(t *testing.T) {
	cleanUp, db := setup(t)
	defer cleanUp("items")

	// test-case setup
	_, err := db.Exec(`INSERT INTO items (name) VALUES ('item 1'), ('item 2')`)
	require.NoError(t, err)

	req, err := http.NewRequest("GET", "http://localhost:8888/items", nil)
	require.NoError(t, err)

	// execution
	httpCli := http.DefaultClient
	resp, err := httpCli.Do(req)
	require.NoError(t, err)

	// validation
	require.Equal(t, http.StatusOK, resp.StatusCode)

	respData, err := io.ReadAll(resp.Body)
	require.NoError(t, err)

	expected := `	
	[
		{
			"id": 1,
			"name": "item 1"
		},
		{
			"id": 2,
			"name": "item 2"
		}
	]`

	assert.JSONEq(t, expected, string(respData))
}
