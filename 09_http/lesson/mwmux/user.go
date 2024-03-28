package main

type User struct {
	Login    string
	Password string
	Token    string
}

var (
	users = []*User{
		&User{
			Login:    "admin",
			Password: "x",
		},
	}
)

func authenticateUser(login, password string) (string, bool) {
	for _, u := range users {
		if u.Login == login && u.Password == password {
			u.Token = "some_random_string"
			return u.Token, true
		}
	}

	return "", false
}

func checkToken(token string) bool {
	for _, u := range users {
		if u.Token == token {
			return true
		}
	}

	return false
}
