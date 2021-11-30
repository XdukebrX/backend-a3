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

	//get all products route
	router.HandleFunc("/listproducts", server.GetAllProducts).Methods("GET")
	//get all raw materials route
	router.HandleFunc("/listmaterials", server.GetAllRawMaterials).Methods("GET")
	//get all commodities route
	router.HandleFunc("/listcommodities", server.GetAllCommodities).Methods("GET")
	//get product by name
	router.HandleFunc("/listproducts/{name}", server.GetProductByName).Methods("GET")

	//starting server
	fmt.Println("Listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))

}
