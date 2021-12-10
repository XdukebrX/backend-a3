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
	Id_product      uint32 `json:"idProduct"`
	Id_raw_material uint32 `json:"idRawmaterial"`
	Quantity        uint32 `json:"quantity"`
	Product         `json:"product"`
	RawMaterial     `json:"rawMaterial"`
}

//create new product
func CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product Product
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
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
	insert, err := db.Query("INSERT INTO products(pname, pvalue) VALUES(?, ?)", product.Name, product.Value)
	if err != nil {
		fmt.Println(err)
	}
	defer insert.Close()
	w.WriteHeader(http.
		StatusOK) // 201

}

//create new raw material
func CreateRawMaterial(w http.ResponseWriter, r *http.Request) {
	var rawMaterial RawMaterial
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
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
	insert, err := db.Query("INSERT INTO raw_materials(rname, stock) VALUES(?, ?)", rawMaterial.Name, rawMaterial.Stock)
	if err != nil {
		fmt.Println(err)
	}
	defer insert.Close()
	w.WriteHeader(http.
		StatusOK) // 201

}

//create new commodity
func CreateCommodity(w http.ResponseWriter, r *http.Request) {
	var commodity Commodity
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	json.Unmarshal(reqBody, &commodity)

	db, err := database.DbConnect()
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	//inserting new commodity
	db.Exec("SET GLOBAL FOREIGN_KEY_CHECKS=0;")
	insert, err := db.Query("INSERT INTO commodities(id_products, id_raw_materials, quantity) VALUES(?, ?, ?)", commodity.Id_product, commodity.Id_raw_material, commodity.Quantity)
	if err != nil {
		fmt.Println(err)
	}
	defer insert.Close()
	w.WriteHeader(http.StatusOK) // 201

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

//get product by id_product via url
func GetProductById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", JSON)
	params := mux.Vars(r)
	id := params["id"]
	db, err := database.DbConnect()
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	//getting product by id_product
	rows, err := db.Query("SELECT * FROM products WHERE id_product = ?", id)
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()

	var product Product
	for rows.Next() {
		err := rows.Scan(&product.Id, &product.Name, &product.Value)
		if err != nil {
			fmt.Println(err)
		}
	}
	json.NewEncoder(w).Encode(product)
}

//get raw material by id via url
func GetRawMaterialById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", JSON)
	params := mux.Vars(r)
	rawMaterialId := params["id"]
	db, err := database.DbConnect()
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	//getting raw material by id
	row := db.QueryRow("SELECT * FROM raw_materials WHERE idraw_material = ? ", rawMaterialId)
	var rawMaterial RawMaterial
	err = row.Scan(&rawMaterial.Id, &rawMaterial.Name, &rawMaterial.Stock)
	if err != nil {
		fmt.Println(err)
	}
	json.NewEncoder(w).Encode(rawMaterial)
}

//get commodity by product id
func GetCommodityByProductId(w http.ResponseWriter, r *http.Request) {
	var commodity Commodity
	w.Header().Set("Content-Type", JSON)
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	json.Unmarshal(reqBody, &commodity)
	params := mux.Vars(r)
	productId := params["id"]
	db, err := database.DbConnect()
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	//getting commodity by product and material id
	rows, err := db.Query("SELECT * FROM commodities WHERE id_products = ? ", productId)
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

//get commodity by product id and show product and raw material data
func GetCommodityByProductIdWithData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", JSON)
	params := mux.Vars(r)
	productId := params["id"]
	db, err := database.DbConnect()
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	//getting commodity by product and material id
	rows, err := db.Query("SELECT	commodities.id_products, commodities.id_raw_materials, commodities.quantity, products.id_product , products.pname, products.pvalue ,raw_materials.idraw_material, raw_materials.rname, raw_materials.stock FROM commodities	JOIN products ON commodities.id_products = products.id_product JOIN raw_materials ON commodities.id_raw_materials = raw_materials.idraw_material where products.id_product = ? and raw_materials.idraw_material = commodities.id_raw_materials", productId)

	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()

	var commodities []Commodity

	for rows.Next() {

		var commodity Commodity

		err := rows.Scan(&commodity.Id_product, &commodity.Id_raw_material, &commodity.Quantity, &commodity.Product.Id, &commodity.Product.Name, &commodity.Product.Value, &commodity.RawMaterial.Id, &commodity.RawMaterial.Name, &commodity.RawMaterial.Stock)

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
	row := db.QueryRow("SELECT * FROM products WHERE id_product = ? ", productId)
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
	db.Exec("SET GLOBAL FOREIGN_KEY_CHECKS=0;")
	_, err = db.Exec("UPDATE products SET pname = ?, pvalue = ? WHERE id_product = ?", updatedProduct.Name, updatedProduct.Value, productId)
	if err != nil {
		fmt.Println(err)
	}
	w.WriteHeader(http.StatusOK)
}

//put raw materia
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
	row := db.QueryRow("SELECT * FROM raw_materials WHERE idraw_material = ? ", rawMaterialId)
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
	db.Exec("SET GLOBAL FOREIGN_KEY_CHECKS=0;")
	_, err = db.Exec("UPDATE raw_materials SET rname = ?, stock = ? WHERE idraw_material = ?", updatedRawMaterial.Name, updatedRawMaterial.Stock, rawMaterialId)
	if err != nil {
		fmt.Println(err)
	}
	w.WriteHeader(http.StatusOK)

}

//put commodity by product id and raw material id via url
func UpdateCommodity(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", JSON)
	params := mux.Vars(r)
	productId := params["id_product"]
	rawMaterialId := params["id_raw_material"]
	db, err := database.DbConnect()
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	//getting commodity by product and material id
	row := db.QueryRow("SELECT * FROM commodities WHERE id_products = ? AND id_raw_materials = ?", productId, rawMaterialId)
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
	db.Exec("SET GLOBAL FOREIGN_KEY_CHECKS=0;")
	_, err = db.Exec("UPDATE commodities SET quantity = ? WHERE id_products = ? AND id_raw_materials = ?", updatedCommodity.Quantity, productId, rawMaterialId)
	if err != nil {
		fmt.Println(err)
	}

	w.WriteHeader(http.StatusOK)
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
}
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
	_, err = db.Exec("DELETE FROM raw_materials WHERE idraw_material = ?", id_product)
	if err != nil {
		fmt.Println(err)
	}
	w.WriteHeader(http.StatusOK)
}

//delete commodity by product id and raw material id via url
func DeleteCommodity(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", JSON)
	params := mux.Vars(r)
	productId := params["id1"]

	db, err := database.DbConnect()
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	//deleting commodity
	_, err = db.Exec("DELETE FROM commodities WHERE id_products = ?", productId)
	if err != nil {
		fmt.Println(err)
	}
	w.WriteHeader(http.StatusOK)
}
