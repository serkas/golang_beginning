package api

import (
	"net/http"
)

func (*API) Hello(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("Hello"))
}
