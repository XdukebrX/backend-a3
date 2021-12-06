package server

import (
	"backend-a3/database"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
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
	Id_raw_material uint32 `json:"id_raw_material"`
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

//get product by name via url
func GetProductByName(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", JSON)
	params := mux.Vars(r)
	productName := params["name"]
	db, err := database.DbConnect()
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	productName = "%" + productName + "%"
	//getting product by name
	row := db.QueryRow("SELECT * FROM products WHERE name LIKE  ? ", productName)
	var product Product
	err = row.Scan(&product.Id, &product.Name, &product.Value)
	if err != nil {
		fmt.Println(err)
	}
	json.NewEncoder(w).Encode(product)
}

//get raw material by name via url
func GetRawMaterialByName(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", JSON)
	params := mux.Vars(r)
	rawMaterialName := params["name"]
	db, err := database.DbConnect()
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	rawMaterialName = "%" + rawMaterialName + "%"
	//getting raw material by name
	row := db.QueryRow("SELECT * FROM raw_materials WHERE name LIKE  ? ", rawMaterialName)
	var rawMaterial RawMaterial
	err = row.Scan(&rawMaterial.Id, &rawMaterial.Name, &rawMaterial.Stock)
	if err != nil {
		fmt.Println(err)
	}
	json.NewEncoder(w).Encode(rawMaterial)
}

//get commodity by product id
func GetCommodityByProductId(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", JSON)
	params := mux.Vars(r)
	productId := params["id"]
	db, err := database.DbConnect()
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	//getting commodity by product id
	rows, err := db.Query("SELECT * FROM commodities WHERE id_product = ?", productId)
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

//put product
func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", JSON)
	params := mux.Vars(r)
	productId := params["id"]
	db, err := database.DbConnect()
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	//getting product by id
	row := db.QueryRow("SELECT * FROM products WHERE id = ? ", productId)
	var product Product
	err = row.Scan(&product.Id, &product.Name, &product.Value)
	if err != nil {
		fmt.Println(err)
	}

	//decoding json
	decoder := json.NewDecoder(r.Body)
	var updatedProduct Product
	err = decoder.Decode(&updatedProduct)
	if err != nil {
		fmt.Println(err)
	}

	//updating product
	_, err = db.Exec("UPDATE products SET name = ?, value = ? WHERE id_product = ?", updatedProduct.Name, updatedProduct.Value, productId)
	if err != nil {
		fmt.Println(err)
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Product has been successfully updated")
}

//put raw material
func UpdateRawMaterial(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", JSON)
	params := mux.Vars(r)
	rawMaterialId := params["id"]
	db, err := database.DbConnect()
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	//getting raw material by id
	row := db.QueryRow("SELECT * FROM raw_materials WHERE id = ? ", rawMaterialId)
	var rawMaterial RawMaterial
	err = row.Scan(&rawMaterial.Id, &rawMaterial.Name, &rawMaterial.Stock)
	if err != nil {
		fmt.Println(err)
	}

	//decoding json
	decoder := json.NewDecoder(r.Body)
	var updatedRawMaterial RawMaterial
	err = decoder.Decode(&updatedRawMaterial)
	if err != nil {
		fmt.Println(err)
	}

	//updating raw material
	_, err = db.Exec("UPDATE raw_materials SET name = ?, stock = ? WHERE id_raw_material = ?", updatedRawMaterial.Name, updatedRawMaterial.Stock, rawMaterialId)
	if err != nil {
		fmt.Println(err)
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Raw material has been successfully updated")
}

//put commodity
func UpdateCommodity(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", JSON)
	params := mux.Vars(r)
	Id_product := params["id"]
	db, err := database.DbConnect()
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	//getting commodity by id
	row := db.QueryRow("SELECT * FROM commodities WHERE id_commodity = ? ", Id_product)
	var commodity Commodity
	err = row.Scan(&commodity.Id_product, &commodity.Id_raw_material, &commodity.Quantity)
	if err != nil {
		fmt.Println(err)
	}

	//decoding json
	decoder := json.NewDecoder(r.Body)
	var updatedCommodity Commodity
	err = decoder.Decode(&updatedCommodity)
	if err != nil {
		fmt.Println(err)
	}

	//updating commodity
	_, err = db.Exec("UPDATE commodities SET id_raw_material = ?, quantity = ?, id_product =? WHERE id_product = ?", updatedCommodity.Id_raw_material, updatedCommodity.Quantity, updatedCommodity.Id_product, Id_product)
	if err != nil {
		fmt.Println(err)
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Commodity has been successfully updated")
}

//delete product
func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", JSON)
	params := mux.Vars(r)
	productId := params["id"]
	db, err := database.DbConnect()
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	//deleting product
	_, err = db.Exec("DELETE FROM products WHERE id_product = ?", productId)
	if err != nil {
		fmt.Println(err)
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Product has been successfully deleted")
}

//delete raw material
func DeleteRawMaterial(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", JSON)
	params := mux.Vars(r)
	id_product := params["id"]
	db, err := database.DbConnect()
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	//deleting raw material
	_, err = db.Exec("DELETE FROM raw_materials WHERE id_productl = ?", id_product)
	if err != nil {
		fmt.Println(err)
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Raw material has been successfully deleted")
}

//delete commodity
func DeleteCommodity(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", JSON)
	params := mux.Vars(r)
	commodityId := params["id"]
	db, err := database.DbConnect()
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	//deleting commodity
	_, err = db.Exec("DELETE FROM commodities WHERE id_commodity = ?", commodityId)
	if err != nil {
		fmt.Println(err)
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Commodity has been successfully deleted")
}
