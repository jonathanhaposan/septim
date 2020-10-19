package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/jonathanhaposan/septim/septim-backend/component/db"
	"github.com/jonathanhaposan/septim/septim-backend/handler"
	"github.com/jonathanhaposan/septim/septim-backend/internal/repository"
	"github.com/jonathanhaposan/septim/septim-backend/router"
	"github.com/jonathanhaposan/septim/septim-backend/server"

	"github.com/julienschmidt/httprouter"
)

var (
	appRouter *httprouter.Router
)

func init() {
	mongoClient := db.NewMongoDB()
	repository := repository.NewRepository(mongoClient)
	handler := handler.Initialize(repository)
	appRouter = router.InitializeRoute(handler)
}

func main() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill, syscall.SIGTERM)

	server.StartWebServer(appRouter)

	<-c

	server.StopWebServer()

	log.Println("Exit")
}
