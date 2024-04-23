package tests

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io"
	"net/http"
	"proj/lessons/16_testing/lesson/functional/app"
	"proj/lessons/16_testing/lesson/functional/model"
	"proj/lessons/16_testing/lesson/functional/storage"
	"testing"
)

func TestAddItem(t *testing.T) {
	// common setup - Usually moved out to a helper function
	store := storage.NewMemStorage() // can be a real DB client if DB is run locally or in Docker
	a := app.New(store)

	defer a.Shutdown()
	go func() {
		err := a.Run(context.Background(), ":8888")
		if err != nil {
			t.Logf("running server: %s", err)
		}
	}()

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
	storedItems, err := store.ListItems(context.Background())
	require.NoError(t, err)
	require.Len(t, storedItems, 1)
	assert.Equal(t, item, storedItems[0])
}

func TestListItems(t *testing.T) {
	// common setup - Usually moved out to a helper function
	store := storage.NewMemStorage() // can be a real DB client if DB is run locally or in Docker
	a := app.New(store)

	defer a.Shutdown()
	go func() {
		err := a.Run(context.Background(), ":8888")
		if err != nil {
			t.Logf("running server: %s", err)
		}
	}()

	// test-case setup
	err := store.AddItem(context.Background(), &model.Item{ID: 1, Name: "item 1"})
	require.NoError(t, err)
	err = store.AddItem(context.Background(), &model.Item{ID: 2, Name: "item 2"})
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
			"ID": 1,
			"Name": "item 1"
		},
		{
			"ID": 2,
			"Name": "item 2"
		}
	]`

	assert.JSONEq(t, expected, string(respData))
}
