package api

import (
	"net/http"
	"time"
)

func (a *API) Hello(w http.ResponseWriter, _ *http.Request) {
	a.log.Info("start handling hello")

	time.Sleep(5 * time.Second)

	a.log.Info("hello response ready")

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"message": "hello"`))
}
