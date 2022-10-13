package project

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Product struct {
	Id          int     `json:"id"`
	ProductName string  `json:"productName"`
	CategoryId  int     `json:"categoryId"`
	UnitPrice   float64 `json:"unitPrice"`
}

type Category struct {
	Id           int    `json:"id"`
	CategoryName string `json:"categoryName"`
}

// func GetAllProducts() {
// 	response, err := http.Get("http://localhost:3000/products")

// 	if err != nil { // Hata var ise
// 		fmt.Println(err)
// 	}
// 	defer response.Body.Close() // Kapatıyoruz.

// 	bodyBytes, _ := ioutil.ReadAll(response.Body)

// 	var products []Product
// 	json.Unmarshal(bodyBytes, &products)
// 	fmt.Println(products)
// }

func GetAllProducts() ([]Product, error) {
	response, err := http.Get("http://localhost:3000/products")

	if err != nil { // Hata var ise
		return nil, err
	}
	defer response.Body.Close() // Kapatıyoruz.

	bodyBytes, _ := ioutil.ReadAll(response.Body)

	var products []Product
	json.Unmarshal(bodyBytes, &products)
	return products, nil
}

func AddProduct() (Product, error) { // Refactoring Yapalım (Yeniden yapılandırma)
	// product := Product{Id: 4, ProductName: "Telephone", CategoryId: 1, UnitPrice: 6000.0}
	product := Product{Id: 9, ProductName: "Microphone 3", CategoryId: 1, UnitPrice: 2000.0}

	jsonProduct, err := json.Marshal(product)

	response, err := http.Post("http://localhost:3000/products", "application/json;charset=utf-8", bytes.NewBuffer(jsonProduct))

	if err != nil { // Hata var ise
		return Product{}, err
	}

	defer response.Body.Close()

	bodyBytes, _ := ioutil.ReadAll(response.Body)

	var productResponse Product
	json.Unmarshal(bodyBytes, &productResponse)

	//fmt.Println("Kaydedildi: ", productResponse)
	return productResponse, nil
}
