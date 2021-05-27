package middlewares

import (
	"net/http"

	"github.com/tobias/todo-app/config"
	res "github.com/tobias/todo-app/utils"
)

func CheckAuthentication(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !config.CheckAuth(w, r) {
			res.JSON(w, 400, "Please log in")
			return
		}
		next.ServeHTTP(w, r)
	}
}
