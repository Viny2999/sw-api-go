package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"net/http"
)

var results []bson.M

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/planet", getAllPlanet).Methods("GET")
	router.HandleFunc("/planet/{id}", getPlanetId).Methods("GET")
	router.HandleFunc("/planet/search/{name}", getPlanetName).Methods("GET")
	router.HandleFunc("/planet/{id}", createPlanet).Methods("POST")
	router.HandleFunc("/planet/{id}", deletePlanet).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", router))
}

func getAllPlanet(w http.ResponseWriter, r *http.Request) {
	results = getAll()
	json.NewEncoder(w).Encode(results)
}

func getPlanetId(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var result = getOneId(params["id"])
	json.NewEncoder(w).Encode(result)
}

func getPlanetName(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var result = getOneName(params["name"])
	json.NewEncoder(w).Encode(result)
}

func createPlanet(w http.ResponseWriter, r *http.Request) {

}

func deletePlanet(w http.ResponseWriter, r *http.Request) {
	//params := mux.Vars(r)
}
