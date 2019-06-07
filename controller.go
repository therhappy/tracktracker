package main

import (
	"errors"
	"log"

	"go.mongodb.org/mongo-driver/bson"
)

func buyProduct(productShortname string, price, unitPrice float32) {

	exists, err := existSN(productShortname)
	checkPanic(err)

	if !exists {
		log.Fatal(errors.New("Product doesn't exist"))
	}

	if unitPrice < 0 {
		log.Fatal(errors.New("Unit price is negative"))
	}

	units := price / unitPrice

	incSpent(productShortname, price)
	incCount(productShortname, units)

}

func goUpdateValues() {
	err := updateAllValues()
	if err != nil {
		log.Fatal(err)
	}
}

func createProductByISIN(isin string) {
	newProduct, err := productInfoByISIN(isin)
	checkPanic(err)
	_, err = insertProduct(newProduct)
	checkPanic(err)
}

func getWorth() float32 {
	var worth float32

	allProducts, err := getProducts(bson.M{})
	checkPanic(err)

	for i := 0; i < len(allProducts); i++ {
		worth += allProducts[i].UnitValue * allProducts[i].UnitCount
	}

	return worth
}
