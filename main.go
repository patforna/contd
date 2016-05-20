package main

import (
	"log"
	"net/http"

	"github.com/patforna/contd/web"
)

func main() {
	log.Fatal(http.ListenAndServe(":8080", web.NewRouter()))
}
