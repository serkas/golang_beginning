package main

import "net/http"

func authorize(nextHandler func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		c, err := r.Cookie("session_token")
		if err != nil {
			if err == http.ErrNoCookie {
				// If the cookie is not set, return an unauthorized status
				http.Redirect(w, r, "/login", http.StatusSeeOther)
				return
			}
			// For any other type of error, return a bad request status
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if checkToken(c.Value) {
			nextHandler(w, r)
			return
		}

		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}
}
