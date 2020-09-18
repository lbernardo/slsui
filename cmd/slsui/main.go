package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/fatih/color"
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
	cyan := color.New(color.FgCyan).SprintFunc()
	fmt.Println("App running at:")
	fmt.Printf("- Local: %v\n", cyan("http://127.0.0.1:8085"))

	log.Println(http.ListenAndServe("0.0.0.0:8085", r))
}
