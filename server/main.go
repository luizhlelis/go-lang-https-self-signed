package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/luizhlelis/go-lang-https-self-signed/server/controllers"
)

func main() {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env files")
	}

	handleRequests()
}

func handleRequests() {

	tlsCert := os.Getenv("tls-certificate")
	tlsKey := os.Getenv("tls-key")
	serverPort := os.Getenv("server-port")

	router := mux.NewRouter().StrictSlash(true)
	controllers.HandleHomeRoutes(router, "https")

	log.Fatal(http.ListenAndServeTLS(serverPort, tlsCert, tlsKey, router))
}
