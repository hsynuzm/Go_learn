package main

import (
	"fmt"
	"reflect"
)

func main() {

	/* first method
	var productName string
	var quantity int
	var discount float32
	var isInStock bool

	productName = "i am learning go"
	quantity = 4
	discount = 0.35
	isInStock = true

	fmt.Println(productName, "-->", reflect.TypeOf(productName))
	fmt.Println(quantity, "-->", reflect.TypeOf(quantity))
	fmt.Println(discount, "-->", reflect.TypeOf(discount))
	fmt.Println(isInStock, "-->", reflect.TypeOf(isInStock)) */

	/* second method
	var productName string = "i ma learning go"
	fmt.Println(productName, reflect.TypeOf(productName)) */

	productName := "i am learning go"
	fmt.Println(productName, reflect.TypeOf(productName))

}
