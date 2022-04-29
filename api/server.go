package main

import (
	"github.com/unrolled/render"
	"go-community/api/service"
	"log"
	"net/http"
)

func main() {
	ip := "127.0.0.1"
	port := "3000"

	r := render.New()
	mux := http.NewServeMux()

	service.Router(r, mux)

	err := http.ListenAndServe(ip+":"+port, mux)
	if err != nil {
		log.Panic(err)
	}
}
