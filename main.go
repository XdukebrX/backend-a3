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

	//starting server
	fmt.Println("Escutando na porta 8080")
	log.Fatal(http.ListenAndServe(":8080", router))

}
