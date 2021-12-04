package main

import (
	"backend-a3/server"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	fmt.Println("Starting the application...")

	//creating new router
	router := mux.NewRouter()
	//adding new route
	router.HandleFunc("/products", server.CreateProduct).Methods("POST")
	//create material
	router.HandleFunc("/materials", server.CreateRawMaterial).Methods("POST")
	//create commodity
	router.HandleFunc("/commodities", server.CreateCommodity).Methods("POST")

	//get all products
	router.HandleFunc("/listproducts", server.GetAllProducts).Methods("GET")
	//get all raw materials
	router.HandleFunc("/listmaterials", server.GetAllRawMaterials).Methods("GET")
	//get all commodities
	router.HandleFunc("/listcommodities", server.GetAllCommodities).Methods("GET")
	//get product by name
	router.HandleFunc("/listproducts/{name}", server.GetProductByName).Methods("GET")
	//get raw material by name
	router.HandleFunc("/listmaterials/{name}", server.GetRawMaterialByName).Methods("GET")
	//get commodity by product id
	router.HandleFunc("/listcommodities/{id}", server.GetCommodityByProductId).Methods("GET")

	//put product
	router.HandleFunc("/alterproducts/{id}", server.UpdateProduct).Methods("PUT")
	//put raw material
	router.HandleFunc("/altermaterials/{id}", server.UpdateRawMaterial).Methods("PUT")
	//put commodity
	router.HandleFunc("/altercommodities/{id}", server.UpdateCommodity).Methods("PUT")

	//starting server

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://*:3000"},
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
	})

	handler := c.Handler(router)

	fmt.Println("Listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", handler))

}
