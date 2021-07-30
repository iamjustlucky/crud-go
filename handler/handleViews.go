package handler

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"text/template"
)

func FrontHomePage(w http.ResponseWriter, r *http.Request){
	fmt.Println("Endpoint hit : Home Page")
	if r.URL.Path != "/"{
		http.NotFound(w, r)
        return
    }
	
	tmpl, err := template.ParseFiles(path.Join("views", "index.html"), path.Join("views", "layout.html"))
	if err != nil{
		log.Println(err)
		http.Error(w, "Error", http.StatusInternalServerError)
		return
	}
	
	err = tmpl.Execute(w, nil)
	if err != nil{
		log.Println(err)
		http.Error(w, "Error", http.StatusInternalServerError)
		return
	}
}

func FrontGetProducts(w http.ResponseWriter, r *http.Request){
	fmt.Println("Endpoint hit : FrontEnd Get Products")

	if r.URL.Path != "/products"{
		http.NotFound(w, r)
        return
    }
	
	tmpl, err := template.ParseFiles(path.Join("views", "products.html"), path.Join("views", "layout.html"))
	if err != nil{
		log.Println(err)
		http.Error(w, "Error", http.StatusInternalServerError)
		return
	}
	
	err = tmpl.Execute(w, nil)
	if err != nil{
		log.Println(err)
		http.Error(w, "Error", http.StatusInternalServerError)
		return
	}

	resp, err := http.Get("http://localhost:9999/api/products/")
    if err != nil {
        fmt.Print(err.Error())
        os.Exit(1)
    }

    if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
		log.Fatalln(err)
	}

	sb := string(body)
	w.Write([]byte(sb))
}