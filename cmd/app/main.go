package main

import (
	"log"

	"github.com/glitcherzigo/book-api/actions"
)

func main() {
	srv := actions.App()

	log.Fatalln(srv.ListenAndServe())
}
