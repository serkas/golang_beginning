package tests_test

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/stretchr/testify/require"
	"go.uber.org/fx"
	"net/http"
	"proj/lessons/21_di/lesson/service/internal/app"
	"proj/lessons/21_di/lesson/service/internal/model"
	"proj/lessons/21_di/lesson/service/internal/services"
	"testing"
	"time"
)

func TestCreate(t *testing.T) {
	stop := setup(t)
	defer stop()

	// test-case setup
	item := &model.Item{
		Name: "test item",
	}
	data, err := json.Marshal(item)
	require.NoError(t, err)

	req, err := http.NewRequest("POST", "http://localhost:8081/items", bytes.NewReader(data))
	require.NoError(t, err)

	// execution
	cli := http.DefaultClient
	resp, err := cli.Do(req)
	require.NoError(t, err)

	// check expectations
	require.Equal(t, http.StatusCreated, resp.StatusCode)

	// more checking of expectations

	// Imagine we need to check all fields are set including timestamp

	// checks are skipped for simplicity
	// what do we want - predictable and stable data in the table
}

func TestCreate2(t *testing.T) {
	stop := setup(t)
	defer stop()

	// test-case setup
	item := &model.Item{
		Name: "test item",
	}
	data, err := json.Marshal(item)
	require.NoError(t, err)

	req, err := http.NewRequest("POST", "http://localhost:8081/items", bytes.NewReader(data))
	require.NoError(t, err)

	// execution
	cli := http.DefaultClient
	resp, err := cli.Do(req)
	require.NoError(t, err)

	// check expectations
	require.Equal(t, http.StatusCreated, resp.StatusCode)

	// more checking of expectations
}

func TestCreate3(t *testing.T) {
	stop := setup(t)
	defer stop()

	// test-case setup
	item := &model.Item{
		Name: "test item",
	}
	data, err := json.Marshal(item)
	require.NoError(t, err)

	req, err := http.NewRequest("POST", "http://localhost:8081/items", bytes.NewReader(data))
	require.NoError(t, err)

	// execution
	cli := http.DefaultClient
	resp, err := cli.Do(req)
	require.NoError(t, err)

	// check expectations
	require.Equal(t, http.StatusCreated, resp.StatusCode)

	// more checking of expectations
}

func setup(t *testing.T) (stop func()) {
	fxApp := fx.New(
		app.New(),
		fx.Replace(newNowFuncForTest()), // caution! replace expects an instance of dependency not a constructor
		fx.Invoke(app.RegisterHTTPServer),
	)

	if err := fxApp.Start(context.Background()); err != nil {
		t.Fatal(err)
	}

	return func() {
		// DB cleanup is skipped
		if err := fxApp.Stop(context.Background()); err != nil {
			t.Error(err)
		}
	}
}

func newNowFuncForTest() services.NowTimeProvider {
	return func() time.Time {
		return time.Date(2001, 2, 3, 16, 17, 18, 0, time.UTC)
	}
}
