package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", mainPageHandler)

	mux.HandleFunc("/login", loginPageHandler)
	mux.HandleFunc("/data", authorize(userDataHandler))

	s := &http.Server{
		Addr:    ":8001",
		Handler: mux,
	}

	if err := s.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Printf("server serve failed: %s", err)
	}
}

func mainPageHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, `<h1>Main Page</h1><p>Info</p><p><a href="/login">Login</a></p><p><a href="/data">User Data</a></p>`)
}

func userDataHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, `secret user data`)
}

func loginPageHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		html := `<body><h4>Login Page</h4>
				<p>
				<form method="POST">
					<label for="login">Login:</label>
					<input name="login" type="text"><br>
					<label for="password">Password:</label>
					<input name="password" type="password"><br>
					<input type="submit" value="Submit">
				</form>
				</p>
			<body>`
		fmt.Fprint(w, html)
		return
	}

	login := r.FormValue("login")
	password := r.FormValue("password")

	token, ok := authenticateUser(login, password)
	if ok {
		http.SetCookie(w, &http.Cookie{
			Name:    "session_token",
			Value:   token,
			Expires: time.Now().Add(time.Minute),
		})

		http.Redirect(w, r, "/", http.StatusSeeOther)
	}

	// not authenticated, show login form again
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}
