package routes

import (
	"crypto-watchlist-api/controllers"
	"net/http"
)

func SetupSampleRoutes() {
	http.HandleFunc("/sample", controllers.GetSample)
}
