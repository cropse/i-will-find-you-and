package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type App struct {
	Router *mux.Router
}

// Run is server run and port sepcific
func (a *App) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, a.Router))
}
func (a *App) initialize() {
	a.Router = mux.NewRouter()
	a.Router.HandleFunc("/v1/ipinfo/{rawIP}", a.getRegion).Methods("GET")
}
func (a *App) getRegion(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	region, _ := getRegionByIP(vars["rawIP"])

	json.NewEncoder(w).Encode(region)
}
