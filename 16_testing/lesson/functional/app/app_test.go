package app

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"proj/lessons/16_testing/lesson/functional/model"
	"proj/lessons/16_testing/lesson/functional/storage"
)

// More white-box functional testing

func TestAPIHandlerAddItem(t *testing.T) {
	// common setup - Usually moved out to a helper function
	store := storage.NewMemStorage()
	app := New(store, ":8080")

	// test-case setup
	item := &model.Item{
		ID:   2,
		Name: "item 2",
	}
	data, err := json.Marshal(item)
	require.NoError(t, err)

	req, err := http.NewRequest("POST", "/items", bytes.NewReader(data))
	require.NoError(t, err)

	// execution
	respRecorder := httptest.NewRecorder()
	app.addItem(respRecorder, req)

	// validation
	require.Equal(t, http.StatusCreated, respRecorder.Code) // validate response
	// the response body can be validated too if present

	// validate state
	storedItems, err := store.ListItems(context.Background())
	require.NoError(t, err)
	require.Len(t, storedItems, 1)
	assert.Equal(t, item, storedItems[0])
}
