package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/mohdrzu/gomicroservice/models"
)

type Product struct {
	l *log.Logger
}

func NewProduct(l *log.Logger) *Product {
	return &Product{l}
}

// func(p *Product) ServeHTTP(w http.ResponseWriter, r *http.Request){
// 	 if r.Method == http.MethodGet {
// 		p.getProduct(w, r)
// 		return
// 	 }
	 
// 	 if r.Method == http.MethodPost {
// 		p.addProduct(w, r)
// 		return
// 	 }

// 	 if r.Method == http.MethodPut {
// 		reg := regexp.MustCompile(`/([0-9]+)`)
// 		g := reg.FindAllStringSubmatch(r.URL.Path, -1)

// 		if len(g) != 1 {
// 			http.Error(w, "Invalid URI", http.StatusBadRequest)
// 			return
// 		}
// 		if len(g[0]) != 2 {
// 			http.Error(w, "Invalid URI", http.StatusBadRequest)
// 			return
// 		}

// 		idString := g[0][1]
// 		id, err := strconv.Atoi(idString)
// 		if err != nil {
// 			http.Error(w, "Invalid URI", http.StatusBadRequest)
// 			return
// 		}

// 		p.updateProduct(id, w, r)
		
// 		return
// 	 }

// 	 w.WriteHeader(http.StatusMethodNotAllowed)
// }

func(p *Product) GetProduct(w http.ResponseWriter, r *http.Request){
	productList := models.GetProducts()
	err := productList.ToJson(w)
	if err != nil {
		http.Error(w, "unable to marshal json", http.StatusInternalServerError)
	}
}

func (p *Product) AddProduct(w http.ResponseWriter, r *http.Request) {
	prod := &models.Product{}
	err := prod.FromJson(r.Body)
	if err != nil {
		http.Error(w, "unable to unmarshall json", http.StatusBadRequest)
	}

	models.AddProduct(prod)

}

func (p *Product) UpdateProduct(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "unable to convert id", http.StatusBadRequest)
		return
	}

	prod := &models.Product{}
	err = prod.FromJson(r.Body)
	if err != nil {
		http.Error(w, "unable to unmarshall json", http.StatusBadRequest)
	}

	err = models.UpdateProduct(id, prod)
	if err == models.ErrProductNotFound {
		http.Error(w, "product not found", http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(w, "product not found", http.StatusNotFound)
		return	
	}
	
	models.UpdateProduct(id, prod)
}