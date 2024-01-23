package handlers

import (
	"log"
	"net/http"

	"github.com/mohdrzu/gomicroservice/models"
)

type Product struct {
	l *log.Logger
}

func NewProduct(l *log.Logger) *Product {
	return &Product{l}
}

func(p *Product) ServeHTTP(w http.ResponseWriter, r *http.Request){
	 if r.Method == http.MethodGet {
		p.getProduct(w, r)
		return
	 }
	 
	 if r.Method == http.MethodPost {
		p.addProduct(w, r)
		return
	 }

	 w.WriteHeader(http.StatusMethodNotAllowed)
}

func(p *Product) getProduct(w http.ResponseWriter, r *http.Request){
	productList := models.GetProducts()
	err := productList.ToJson(w)
	if err != nil {
		http.Error(w, "unable to marshal json", http.StatusInternalServerError)
	}
}

func (p *Product) addProduct(w http.ResponseWriter, r *http.Request) {
	prod := &models.Product{}
	err := prod.FromJson(r.Body)
	if err != nil {
		http.Error(w, "unable to unmarshall json", http.StatusBadRequest)
	}

	models.AddProduct(prod)

}