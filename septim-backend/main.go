package main

import (
	"log"

	"github.com/jonathanhaposan/septim/septim-backend/router"
)

func main() {
	router.InitializeRoute()

	log.Println("Hello World")
}
