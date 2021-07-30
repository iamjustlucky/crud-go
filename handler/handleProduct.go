package handler

import (
	"crud-go/model"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

func GetProducts(w http.ResponseWriter, r *http.Request){
	fmt.Println("Endpoint hit : Get Products")

	products := []model.Product{}
	DB.Find(&products)

	res := model.Result{Code: 200, Data: products, Message: "Success Get Products"}
	result, err := json.Marshal(res)

	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func GetProduct(w http.ResponseWriter, r *http.Request){
	fmt.Println("Endpoint hit: Get Product by ID")
	vars := mux.Vars(r)
	productId := vars["id"]

	var product model.Product
	
	DB.First(&product, productId)

	res := model.Result{Code: 200, Data: product, Message: "Success Get Product by ID"}
	result, err := json.Marshal(res)

	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func InsertProduct(w http.ResponseWriter, r *http.Request){
	fmt.Println("Endpoint hit: insert product")
	payloads, _ := ioutil.ReadAll(r.Body)

	var product model.Product
	json.Unmarshal(payloads, &product)

	DB.Create(&product)

	res := model.Result{Code: 200, Data: product, Message: "Success create product"}
	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func UpdateProduct(w http.ResponseWriter, r *http.Request){
	fmt.Println("Endpoint hit: update product")
	vars := mux.Vars(r)
	productID := vars["id"]

	payloads, _ := ioutil.ReadAll(r.Body)

	var productUpdates model.Product
	json.Unmarshal(payloads, &productUpdates)

	var product model.Product
	DB.First(&product, productID)
	DB.Model(&product).Updates(productUpdates)

	res := model.Result{Code: 200, Data: product, Message: "Success update product"}
	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func DeleteProduct(w http.ResponseWriter, r *http.Request){
	fmt.Println("Endpoint hit: delete product")
	vars := mux.Vars(r)
	productID := vars["id"]

	var product model.Product

	DB.First(&product, productID)
	DB.Delete(&product)

	res := model.Result{Code: 200, Message: "Success delete product"}
	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}