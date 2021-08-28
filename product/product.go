package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)


type Product struct {
	Uuid string `json:"uuid"`
	Product string `json:"product"`
	Price float64 `json:"price,string"`
}


type Products struct {
	Products []Product
}


func loadData() []byte {
	jsonFile, err := os.Open("products.json")
	if err != nil {
		fmt.Println(err.Error())
	}
	defer jsonFile.Close()

	data, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Println(err.Error())
	}

	return data
}


func ListProducts(w http.ResponseWriter, r *http.Request) {
	products := loadData()
	w.Write([]byte(products))
}


func GetProductById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	data := loadData()

	// Interface
	var products Products
	// Desserializa JSON para objeto
	json.Unmarshal(data, &products)
	
	for _, v := range products.Products {
		if v.Uuid == vars["id"] {
			// Serializa objeto para JSON
			product, _ := json.Marshal(v) 
			w.Write([]byte(product))
		}
	}
}


func main() {
	r := mux.NewRouter()
	r.HandleFunc("/products", ListProducts)
	r.HandleFunc("/products/{id}", GetProductById)
	http.ListenAndServe(":8081", r)
}
