package handler

import (
	"crud-go/model"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

func GetContacts(w http.ResponseWriter, r *http.Request){
	fmt.Println("Endpoint hit : get contacts")

	contacts := []model.Contact{}
	DB.Find(&contacts)

	res := model.Result{Code: 200, Data: contacts, Message: "Succes Get Contacts"}
	results, err := json.Marshal(res)

	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(results)
}

func GetContact(w http.ResponseWriter, r *http.Request){
	fmt.Println("Endpoint hit : get contact by id")

	vars := mux.Vars(r)
	productId := vars["id"]

	var contacts model.Contact

	DB.First(contacts, productId)

	res := model.Result{Code: 200, Data: contacts, Message: "Success Get Contact by ID"}
	result, err := json.Marshal(res)

	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func InsertContact(w http.ResponseWriter, r *http.Request){
	fmt.Println("Endpoint hit : insert contact")
	payloads, _ := ioutil.ReadAll(r.Body)

	var contacts model.Contact
	json.Unmarshal(payloads, &contacts)

	DB.Create(&contacts)

	res := model.Result{Code: 200, Data: contacts, Message: "Success create product"}
	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func UpdateContact(w http.ResponseWriter, r *http.Request){
	fmt.Println("Endpoint hit : update contact")

	vars := mux.Vars(r)
	contactId := vars["id"]

	payloads, _ := ioutil.ReadAll(r.Body)
	
	var contactsUpdate model.Contact
	json.Unmarshal(payloads, &contactsUpdate)
	
	var contacts model.Contact 

	DB.First(&contacts, contactId)
	DB.Model(&contacts).Update(contactsUpdate)

	res := model.Result{Code: 200, Data: contacts, Message: "Success update product"}
	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func DeleteContact(w http.ResponseWriter, r *http.Request){
	fmt.Println("Endpoint hit: delete contact")
	vars := mux.Vars(r)
	contactId := vars["id"]

	var contacts model.Contact
	DB.First(&contacts, contactId)
	DB.Delete(&contacts)

	res := model.Result{Code: 200, Data: contacts, Message: "Success delete product"}
	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}