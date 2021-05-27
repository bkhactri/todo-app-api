package config

import (
	"net/http"

	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte("todo-app-secret"))

func Store(w http.ResponseWriter, r *http.Request, email string) {
	session, _ := store.Get(r, "session")
	session.Values["email"] = email
	session.Save(r, w)
}

func CheckAuth(w http.ResponseWriter, r *http.Request) bool {
	session, _ := store.Get(r, "session")
	_, isLogin := session.Values["email"]
	if !isLogin {
		return false
	} else {
		return true
	}
}
