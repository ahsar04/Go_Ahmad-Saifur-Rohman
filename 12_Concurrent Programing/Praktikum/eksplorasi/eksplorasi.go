package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Product struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Price       float64 `json:"price"`
	Category 	string  `json:"category"`
	Description string  `json:"description"`
}

func gatData(url string, ch chan<- []Product) {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	var products []Product
	err = json.NewDecoder(resp.Body).Decode(&products)
	if err != nil {
		panic(err)
	}

	ch <- products
}

func main() {
	baseUrl := "https://fakestoreapi.com/"
	productUrl:="products"
	url:=baseUrl+productUrl
	ch := make(chan []Product)

	go gatData(url, ch)

	products := <-ch
	println("Product Data:")
	for _, product := range products {
		fmt.Println("===")
		fmt.Println("Title: ", product.Title)
		fmt.Println("Price: ", product.Price)
		fmt.Println("Category: ", product.Category)
		// fmt.Println("Description: ", product.Description)
		fmt.Println("===")
	}
}
