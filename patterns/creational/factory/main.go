package main

import "fmt"


type Product struct {
    branch string
    pType string
}


type Factory interface {
    CreateProduct() Product
}


type CarFactory struct{}
func (cf CarFactory) CreateProduct() Product {
    return Product{"Honda", "car"}
}

type BikeFactory struct{}
func (bf BikeFactory) CreateProduct() Product {
    return Product{"Yamaha", "bike"}
}


func getProduct(productType string) Product {
    var factory Factory
    switch productType {
        case "car":
            factory = CarFactory{}
            return factory.CreateProduct()
        case "bike":
            factory = BikeFactory{}
            return factory.CreateProduct()
    }
    return Product{}
}


func main() {
	var productType string
    fmt.Scan(&productType)
    a := getProduct(productType)

    fmt.Println(a.branch)
}
