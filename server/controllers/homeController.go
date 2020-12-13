package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// HandleHomeRoutes is a method to load all home controller routes
func HandleHomeRoutes(router *mux.Router, scheme string) {
	router.HandleFunc("/home", getHome).Schemes(scheme).Methods("GET")
}

func getHome(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Home page endepoint hit")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("This message has been encrypted to get here")
}
