package main

import (
	"log"
	"net/http"

	"github.com/Altyn119/BayernTeam/handlers"

	"github.com/gorilla/mux"
)

func main() {
	log.Println("Starting API server")
	router := mux.NewRouter()

	log.Println("Creating routes")
	router.HandleFunc("/players", handlers.GetPlayers).Methods("GET")
	router.HandleFunc("/players/last_name/{last_name}", handlers.GetName).Methods("GET")
	router.HandleFunc("/players/{id}", handlers.GetID).Methods("GET")
	router.HandleFunc("/players/number/{number}", handlers.GetNumber).Methods("GET")
	router.HandleFunc("/health-checkout", handlers.HealthCheck).Methods("GET")
	router.HandleFunc("/players/position/{position}", handlers.GetPosition).Methods("GET")
	router.HandleFunc("/players/nation/{nation}", handlers.GetNation).Methods("GET")
	http.Handle("/", router)

	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
