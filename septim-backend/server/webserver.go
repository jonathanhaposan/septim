package server

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
)

var server *http.Server

func StartWebServer(router *httprouter.Router) {

	server = &http.Server{
		Addr:         ":8082",
		IdleTimeout:  5 * time.Second,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
		Handler:      router,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Fatalln("failed starting web server", err)
			return
		}
	}()

	log.Println("listen and serve", server.Addr)
}

func StopWebServer() {
	ctx, cancel := context.WithTimeout(context.Background(), 7*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalln("failed shutdown web server properly", err)
		return
	}
}
