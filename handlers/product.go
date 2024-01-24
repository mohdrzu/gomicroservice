package handlers

import (
	"context"
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

func(p *Product) GetProduct(w http.ResponseWriter, r *http.Request){
	productList := models.GetProducts()
	err := productList.ToJson(w)
	if err != nil {
		http.Error(w, "unable to marshal json", http.StatusInternalServerError)
	}
}

func (p *Product) AddProduct(w http.ResponseWriter, r *http.Request) {
	prod := r.Context().Value(KeyProduct{}).(models.Product)
	models.AddProduct(&prod)
}

func (p *Product) UpdateProduct(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "unable to convert id", http.StatusBadRequest)
		return
	}

	prod := r.Context().Value(KeyProduct{}).(models.Product)
	err = models.UpdateProduct(id, &prod)
	if err == models.ErrProductNotFound {
		http.Error(w, "product not found", http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(w, "product not found", http.StatusNotFound)
		return	
	}
}

type KeyProduct struct{}

func(p *Product) MiddlewareValidateProduct(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		prod := &models.Product{}
		
		err := prod.FromJson(r.Body)
		if err != nil {
			p.l.Println("[ERROR] -> deserializing product")
			http.Error(w, "error reading product", http.StatusBadRequest)
			return
		}

		ctx := context.WithValue(r.Context(), KeyProduct{}, prod)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	
	})
}