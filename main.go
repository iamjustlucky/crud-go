package main

import (
	"crud-go/handler"
	"crud-go/model"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main(){
	handler.HandleDB()
	log.Println("Start the development server at http://localhost:9999")

	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)

		res := model.Result{
			Code: 404,
			Message: "Method Not Found",
		}
		response, _ := json.Marshal(res)
		w.Write(response)
	})

	myRouter.MethodNotAllowedHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "applcation/json")
		w.WriteHeader(http.StatusMethodNotAllowed)

		res := model.Result{
			Code: 403,
			Message: "Method Not Allowed",
		}
		response, _ := json.Marshal(res)
		w.Write(response)
	})
	
	//handler product get data to FrontEnd from API
		myRouter.HandleFunc("/", handler.FrontHomePage)
		myRouter.HandleFunc("/products", handler.FrontGetProducts).Methods("GET")

	//handler product DB-API
		//get all product
		myRouter.HandleFunc("/api/products", handler.GetProducts).Methods("GET")
		//get product by id
		myRouter.HandleFunc("/api/products/{id}", handler.GetProduct).Methods("GET")
		//insert products
		myRouter.HandleFunc("/api/products", handler.InsertProduct).Methods("POST")
		//update product by id
		myRouter.HandleFunc("/api/products/{id}", handler.UpdateProduct).Methods("PUT")
		//delete product by id
		myRouter.HandleFunc("/api/products/{id}", handler.DeleteProduct).Methods("DELETE")

	//handler contact
		//get all contact
		myRouter.HandleFunc("/api/contacts", handler.GetContacts).Methods("GET")
		//get product by id
		myRouter.HandleFunc("/api/contacts/{id}", handler.GetContact).Methods("GET")
		//insert new contact
		myRouter.HandleFunc("/api/contacts", handler.InsertContact).Methods("POST")
		//insert new contact
		myRouter.HandleFunc("/api/contacts/{id}", handler.UpdateContact).Methods("PUT")
		//delete contact
		myRouter.HandleFunc("/api/contacts/{id}", handler.DeleteContact).Methods("DELETE")


	log.Fatal(http.ListenAndServe(":9999", myRouter))
}