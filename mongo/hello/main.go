package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
)

/*
// https://www.mongodb.com/docs/drivers/go/current/quick-start/#add-mongodb-as-a-dependency
go get go.mongodb.org/mongo-driver/mongo
go get github.com/joho/godotenv
*/

const conn = "mongodb+srv://dbUser:dbPassword@cluster0.lugjz6h.mongodb.net/?retryWrites=true&w=majority"

func main() {
	ctx := context.TODO()
	opts := options.Client().ApplyURI(conn)

	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		panic(err)
	}

	defer client.Disconnect(ctx)
	fmt.Printf("pinging %T ...\n", client)
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		panic(err)
	}
	fmt.Printf("OK\n")
	dbNames, err := client.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		panic(err)
	}
	fmt.Println(dbNames)

	testDB := client.Database("test")
	fmt.Printf("%T\n", testDB)

	exampleCollection := testDB.Collection("example")
	defer exampleCollection.Drop(ctx)

	fmt.Printf("%T\n", exampleCollection)
	// insert
	/*
	 */
	example := bson.D{
		{"someString", "Example String"},
		{"someInteger", 12},
		{"someStringSlice", []string{"Example 1", "Example 2", "Example 3"}},
	}
	r, err := exampleCollection.InsertOne(ctx, example)
	if err != nil {
		panic(err)
	}
	fmt.Println(r.InsertedID)
	examples := []interface{}{
		bson.D{
			{"someString", "Second Example String"},
			{"someInteger", 253},
			{"someStringSlice", []string{"Example 15", "Example 42", "Example 83", "Example 5"}},
		},
		bson.D{
			{"someString", "Another Example String"},
			{"someInteger", 54},
			{"someStringSlice", []string{"Example 21", "Example 53"}},
		},
	}
	rs, err := exampleCollection.InsertMany(ctx, examples)
	if err != nil {
		panic(err)
	}
	fmt.Println(rs.InsertedIDs)
	time.Sleep(10 * time.Second)

	// query

	c := exampleCollection.FindOne(ctx, bson.M{"_id": r.InsertedID})
	var exampleResult bson.M
	c.Decode(&exampleResult)

	fmt.Printf("\nItem with ID: %v contains the following:\n", exampleResult["_id"])
	fmt.Println("someString:", exampleResult["someString"])
	fmt.Println("someInteger:", exampleResult["someInteger"])
	fmt.Println("someStringSlice:", exampleResult["someStringSlice"])

	filter := bson.D{{"someInteger", bson.D{{"$lt", 60}}}}
	examplesGT50, err := exampleCollection.Find(ctx, filter)
	if err != nil {
		panic(err)
	}
	var examplesResult []bson.M
	if err = examplesGT50.All(ctx, &examplesResult); err != nil {
		panic(err)
	}

	for _, e := range examplesResult {
		fmt.Printf("\nItem with ID: %v contains the following:\n", e["_id"])
		fmt.Println("someString:", e["someString"])
		fmt.Println("someInteger:", e["someInteger"])
		fmt.Println("someStringSlice:", e["someStringSlice"])
	}

	time.Sleep(10 * time.Second)

	// upadate
	rUpdt, err := exampleCollection.UpdateByID(
		ctx,
		r.InsertedID,
		bson.D{
			{"$set", bson.M{"someInteger": 201}},
		},
	)
	if err != nil {
		panic(err)
	}
	fmt.Println("Number of items updated:", rUpdt.ModifiedCount)

	rUpdt, err = exampleCollection.UpdateOne(
		ctx,
		bson.M{"_id": r.InsertedID},
		bson.D{
			{"$set", bson.M{"someString": "The Updated String"}},
		},
	)
	if err != nil {
		panic(err)
	}
	fmt.Println("Number of items updated:", rUpdt.ModifiedCount)

	rUpdt2, err := exampleCollection.UpdateMany(
		ctx,
		bson.D{{"someInteger", bson.D{{"$gt", 60}}}},
		bson.D{
			{"$set", bson.M{"someInteger": 60}},
		},
	)
	if err != nil {
		panic(err)
	}
	fmt.Println("Number of items updated:", rUpdt2.ModifiedCount)

	rDel, err := exampleCollection.DeleteOne(ctx, bson.M{"_id": r.InsertedID})
	if err != nil {
		panic(err)
	}

	fmt.Println("Number of items deleted:", rDel.DeletedCount)
}
