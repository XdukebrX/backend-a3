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
	//get product by id
	router.HandleFunc("/listproducts/{id}", server.GetProductById).Methods("GET")
	//get raw material by id
	router.HandleFunc("/listmaterials/{id}", server.GetRawMaterialById).Methods("GET")
	//get commodity by product id
	router.HandleFunc("/listcommodities/{id}", server.GetCommodityByProductIdWithData).Methods("GET")

	//put product
	router.HandleFunc("/alterproducts/{id}", server.UpdateProduct).Methods("PUT")
	//put raw material
	router.HandleFunc("/altermaterials/{id}", server.UpdateRawMaterial).Methods("PUT")
	//put commodity using product id and commodity id via url
	router.HandleFunc("/altercommodities/{id}/{id2}", server.UpdateCommodity).Methods("PUT")

	//delete product
	router.HandleFunc("/deleteproducts/{id}", server.DeleteProduct).Methods("DELETE")
	//delete raw material
	router.HandleFunc("/deletematerials/{id}", server.DeleteRawMaterial).Methods("DELETE")
	//delete commodity using product id and commodity id via url
	router.HandleFunc("/deletecommodities/{id}", server.DeleteCommodity).Methods("DELETE")

	//starting server

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"*"},
	})

	handler := c.Handler(router)

	fmt.Println("Listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", handler))

}
