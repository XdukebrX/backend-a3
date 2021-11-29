package server

import (
	"backend-a3/database"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const JSON = "application/json"

//struct protuct id name value
type Product struct {
	Id    uint32  `json:"id"`
	Name  string  `json:"name"`
	Value float64 `json:"value"`
}

//struct raw material id name stock
type RawMaterial struct {
	Id    uint32 `json:"id"`
	Name  string `json:"name"`
	Stock uint32 `json:"stock"`
}

//struct commodities  Id_product Id_raw_material quantity
type Commodity struct {
	Id_product      uint32 `json:"id_product"`
	Id_raw_material string `json:"id_raw_material"`
	Quantity        uint32 `json:"quantity"`
}

//create new product
func CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product Product
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the product name and value only in order to update")
		return
	}
	json.Unmarshal(reqBody, &product)
	//fmt.Println(product)
	db, err := database.DbConnect()
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	//inserting new product
	insert, err := db.Query("INSERT INTO products(name, value) VALUES(?, ?)", product.Name, product.Value)
	if err != nil {
		fmt.Println(err)
	}
	defer insert.Close()
	w.WriteHeader(http.
		StatusCreated) // 201

	fmt.Fprintf(w, "New product has been successfully created")
}

//create new raw material
func CreateRawMaterial(w http.ResponseWriter, r *http.Request) {
	var rawMaterial RawMaterial
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the raw material name and stock only in order to update")
		return
	}
	json.Unmarshal(reqBody, &rawMaterial)
	//fmt.Println(rawMaterial)
	db, err := database.DbConnect()
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	//inserting new raw material
	insert, err := db.Query("INSERT INTO raw_materials(name, stock) VALUES(?, ?)", rawMaterial.Name, rawMaterial.Stock)
	if err != nil {
		fmt.Println(err)
	}
	defer insert.Close()
	w.WriteHeader(http.
		StatusCreated) // 201

	fmt.Fprintf(w, "New raw material has been successfully created")
}

//create new commodity
func CreateCommodity(w http.ResponseWriter, r *http.Request) {
	var commodity Commodity
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the product id and raw materials id only in order to update")
		return
	}
	json.Unmarshal(reqBody, &commodity)

	db, err := database.DbConnect()
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	//inserting new commodity
	insert, err := db.Query("INSERT INTO commodities(id_product, id_raw_material, quantity) VALUES(?, ?, ?)", commodity.Id_product, commodity.Id_raw_material, commodity.Quantity)
	if err != nil {
		fmt.Println(err)
	}
	defer insert.Close()
	w.WriteHeader(http.
		StatusCreated) // 201

	fmt.Fprintf(w, "New commodity has been successfully created")

}

//get all products
func GetAllProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", JSON)
	db, err := database.DbConnect()
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	//getting all products
	rows, err := db.Query("SELECT * FROM products")
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()

	var products []Product
	for rows.Next() {
		var product Product
		err := rows.Scan(&product.Id, &product.Name, &product.Value)
		if err != nil {
			fmt.Println(err)
		}
		products = append(products, product)
	}
	json.NewEncoder(w).Encode(products)
}

//get all raw materials
func GetAllRawMaterials(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", JSON)
	db, err := database.DbConnect()
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	//getting all raw materials
	rows, err := db.Query("SELECT * FROM raw_materials")
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()

	var rawMaterials []RawMaterial
	for rows.Next() {
		var rawMaterial RawMaterial
		err := rows.Scan(&rawMaterial.Id, &rawMaterial.Name, &rawMaterial.Stock)
		if err != nil {
			fmt.Println(err)
		}
		rawMaterials = append(rawMaterials, rawMaterial)
	}
	json.NewEncoder(w).Encode(rawMaterials)
}

//get all commodities
func GetAllCommodities(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", JSON)
	db, err := database.DbConnect()
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	//getting all commodities
	rows, err := db.Query("SELECT * FROM commodities")
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()

	var commodities []Commodity
	for rows.Next() {
		var commodity Commodity
		err := rows.Scan(&commodity.Id_product, &commodity.Id_raw_material, &commodity.Quantity)
		if err != nil {
			fmt.Println(err)
		}
		commodities = append(commodities, commodity)
	}
	json.NewEncoder(w).Encode(commodities)
}
