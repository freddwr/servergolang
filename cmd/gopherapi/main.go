package main

import (
	"log"
	"net/http"

	"github.com/freddwr/GO/pkg/server"
)

func main() {
	s := server.New()
	log.Fatal(http.ListenAndServe(":3000", nil))
}
