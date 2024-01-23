package models

import (
	"encoding/json"
	"io"
	"time"
)

type Product struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
	SKU string `json:"sku"`
	Price float64 `json:"price"`
	CreateAt time.Time `json:"-"`
	ModifiedAt time.Time `json:"-"`
	DeletedAt time.Time `json:"-"`
}

func(p *Product) FromJson(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(p)
}


type Products []*Product

func(p *Products) ToJson(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}


func GetProducts() Products {
	return productList
}

func AddProduct(p *Product) {
	p.ID = getNextID()
	productList = append(productList, p)
} 

func getNextID() int {
	lp := productList[len(productList) - 1]
	return lp.ID + 1
}

var productList = []*Product{
	{
		ID: 1,
		Name: "Latte",
		Description: "Frothy milky coffee",
		SKU: "aaa-aaa-aaa-aaa",
		Price: 2.45,
		CreateAt: time.Now().UTC(),
		ModifiedAt:time.Now().UTC(),
	},
	{
		ID: 2,
		Name: "Espresso",
		Description: "Short and strong coffee without milk",
		SKU: "bbb-bbb-bbb-bbb",
		Price: 1.45,
		CreateAt: time.Now().UTC(),
		ModifiedAt:time.Now().UTC(),
	},
}