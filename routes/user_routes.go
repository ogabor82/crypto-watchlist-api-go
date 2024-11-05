package routes

import (
	"crypto-watchlist-api/controllers"
	"net/http"
)

func SetupUserRoutes() {
	http.HandleFunc("/users/signup", controllers.SignupHandler)
	http.HandleFunc("/users/login", controllers.LoginHandler)
}
