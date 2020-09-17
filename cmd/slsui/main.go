package main

import (
	"log"
	"net/http"

	"github.com/gobuffalo/packr"
	"github.com/gorilla/mux"
	"github.com/lbernardo/slsui/internal/primary"
)

func main() {
	box := packr.NewBox("../../webui")
	r := mux.NewRouter()
	service := primary.NewService()

	r.HandleFunc("/api/build", service.Build).Methods("POST")
	r.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(box)))

	log.Println(http.ListenAndServe("0.0.0.0:8085", r))
}
