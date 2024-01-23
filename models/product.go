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

type Products []*Product

func(p *Products) ToJson(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func GetProducts() Products {
	return ProductList
}

var ProductList = []*Product{
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