package main

import (
	"backend-a3/server"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Starting the application...")
	//creating new router
	router := mux.NewRouter()
	//adding new route
	router.HandleFunc("/products", server.CreateProduct).Methods("POST")
	//create material route
	router.HandleFunc("/materials", server.CreateRawMaterial).Methods("POST")
	//create commodity route
	router.HandleFunc("/commodities", server.CreateCommodity).Methods("POST")
	//create get all products route
	router.HandleFunc("/products", server.GetAllProducts).Methods("GET")
	//create get all raw materials route
	router.HandleFunc("/materials", server.GetAllRawMaterials).Methods("GET")
	//create get all commodities route
	router.HandleFunc("/commodities", server.GetAllCommodities).Methods("GET")
	//starting server
	fmt.Println("Escutando na porta 8080")
	log.Fatal(http.ListenAndServe(":8080", router))

}
