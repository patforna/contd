package main

import (
	"log"
	"net/http"

	"github.com/patforna/splendid/web"
)

func main() {
	log.Fatal(http.ListenAndServe(":8080", web.NewRouter()))
}