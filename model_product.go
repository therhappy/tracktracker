package main

import (
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Product : Generic financial product structure
type Product struct {
	Name         string  `json:"name" bson:"name"`
	ShortName    string  `json:"shortName" bson:"shortName"`
	ISIN         string  `json:"isin" bson:"isin"`
	UnitValue    float32 `json:"unitValue" bson:"unitValue"`
	UnitCount    float32 `json:"unitCost" bson:"unitCount"`
	TotalBuyCost float32 `json:"totalBuyCost" bson:"totalBuyCost"`
	TotalFees    float32 `json:"totalFees" bson:"totalFees"`
}

func insertDummy() {

	// Some dummy data to add to the Database
	ex1 := Product{"Amundi Test", "AT", "FR01234556", 10.05, 23.45, 213.5, 1.1}
	ex2 := Product{"Bourso Test", "BT", "LU01223442", 5.10, 12, 80, 1.0}
	ex3 := Product{"Clodoa Test", "CT", "US12345678", 23, 12, 34, 2}

	// Insert a single document
	insertResult, err := dbCollection.InsertOne(ctx, ex1)
	checkPanic(err)

	fmt.Println("Inserted a single document: ", insertResult.InsertedID)

	// Insert multiple documents
	products := []interface{}{ex2, ex3}

	insertManyResult, err := dbCollection.InsertMany(ctx, products)
	checkPanic(err)
	fmt.Println("Inserted multiple documents: ", insertManyResult.InsertedIDs)
}

func insertProduct(prod Product) (*mongo.InsertOneResult, error) {
	return dbCollection.InsertOne(ctx, prod)
}

func incCount(productShortname string, difference float32) (*mongo.UpdateResult, error) {
	filter := bson.M{"shortName": productShortname}

	update := bson.D{
		{"$set", bson.D{
			{"UnitCount", difference},
		}},
	}

	return dbCollection.UpdateOne(ctx, filter, update)
}

func changeUnitValue(productISIN string, newValue float32) (*mongo.UpdateResult, error) {
	filter := bson.M{"ISIN": productISIN} // recommended. May generalize.
	update := bson.D{
		{"$set", bson.D{
			{"UnitValue", newValue},
		}},
	}

	return dbCollection.UpdateOne(ctx, filter, update)
}

func incSpent(productShortname string, difference float32) (*mongo.UpdateResult, error) {
	filter := bson.M{"shortName": productShortname}

	update := bson.D{
		{"$inc", bson.D{
			{"UnitValue", difference},
		}},
	}

	return dbCollection.UpdateOne(ctx, filter, update)
}

// getSNlist : returns an array of all Products ShortName in the database
func getSNlist() ([]*string, error) {
	findOptions := options.Find()
	cur, err := dbCollection.Find(ctx, bson.D{{}}, findOptions)
	var results []*string

	// Iterate
	for cur.Next(ctx) {
		var elem Product
		err := cur.Decode(&elem)
		checkPanic(err)
		results = append(results, &elem.ShortName)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	cur.Close(ctx)
	return results, err
}

// getISINlist : returns an array of all Products ISIN in the database
func getISINlist() ([]*string, error) {
	findOptions := options.Find()
	cur, err := dbCollection.Find(ctx, bson.M{}, findOptions)
	var results []*string

	// Iterate
	for cur.Next(ctx) {
		var elem Product
		err := cur.Decode(&elem)
		checkPanic(err)
		results = append(results, &elem.ShortName)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	cur.Close(ctx)
	return results, err
}

func updateAllValues() error {
	findOptions := options.Find()
	cur, err := dbCollection.Find(ctx, bson.D{{}}, findOptions)
	if err != nil {
		return err
	}

	// Iterate
	for cur.Next(ctx) {
		var elem Product
		err := cur.Decode(&elem)
		if err != nil {
			return err
		}

		_, err = changeUnitValue(elem.ISIN, extAPIgetValue(elem.ISIN))
		if err != nil {
			return err
		}
	}
	cur.Close(ctx)
	return nil
}

func existSN(prodSN string) (bool, error) {
	filter := bson.M{"shortName": prodSN}
	n, err := dbCollection.CountDocuments(ctx, filter, nil) // nil could work for CountOptions

	if err != nil {
		return false, err
	}

	if n != 0 {
		return true, nil
	}
	return false, nil
}

func getProducts(filter map[string]interface{}) ([]*Product, error) {
	findOptions := options.Find()
	cur, err := dbCollection.Find(ctx, filter, findOptions)
	var results []*Product

	// Iterate
	for cur.Next(ctx) {
		var elem Product
		err := cur.Decode(&elem)
		checkPanic(err)
		results = append(results, &elem)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	cur.Close(ctx)
	return results, err
}
