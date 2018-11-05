package main

import (
	"fmt"
	"gbmchallenge/api/route"
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

func main() {
	r := mux.NewRouter()
	route.CreateRoutes(r)

	server := &http.Server{
		Addr:           ":5000",
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	fmt.Println("Listening on port 5000")
	server.ListenAndServe()
}
