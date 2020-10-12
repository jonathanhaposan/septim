package router

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

var r *httprouter.Router

func InitializeRoute() *httprouter.Router {
	r = httprouter.New()

	allowCORS()
	registerRoutes()

	return r
}

func registerRoutes() {
	r.GET("/purchase", nil)
	r.GET("/purchase/:id", nil)
	r.PUT("/purchase/:id", nil)
	r.POST("/purchase", nil)
}

func allowCORS() {
	r.GlobalOPTIONS = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Access-Control-Request-Method") != "" {
			// Set CORS headers
			header := w.Header()
			header.Set("Access-Control-Allow-Methods", header.Get("Allow"))
			header.Set("Access-Control-Allow-Origin", "*")
		}

		w.WriteHeader(http.StatusNoContent)
	})
}
