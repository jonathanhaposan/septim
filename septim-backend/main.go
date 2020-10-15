package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/jonathanhaposan/septim/septim-backend/router"
	"github.com/jonathanhaposan/septim/septim-backend/server"

	"github.com/julienschmidt/httprouter"
)

var (
	appRouter *httprouter.Router
)

func init() {
	appRouter = router.InitializeRoute()
}

func main() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill, syscall.SIGTERM)

	server.StartWebServer(appRouter)

	<-c

	server.StopWebServer()

	log.Println("Exit")
}
