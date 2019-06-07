package main

import (
	"log"
	"time"
)

func checkPanic(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {

	// initDB : instanciates global vars Client, CurrentCollection
	initDB()

	// insert documents
	insertDummy()

	// update a product quantity
	buyProduct("AT", 25, 10.05)

	// Sum worth of products
	worth := getWorth()
	println("Current worth is", worth)

	// Wait a sec
	time.Sleep(2 * time.Second)

	// Delete all the documents in the collection
	eraseDB(true)

	disconnectDB()
}
