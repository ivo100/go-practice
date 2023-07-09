package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

/*
// https://www.mongodb.com/docs/drivers/go/current/quick-start/#add-mongodb-as-a-dependency
go get go.mongodb.org/mongo-driver/mongo
go get github.com/joho/godotenv
*/

const conn = "mongodb+srv://dbUser:dbPassword@cluster0.lugjz6h.mongodb.net/?retryWrites=true&w=majority"

var client *mongo.Client
var coll *mongo.Collection
var id primitive.ObjectID

func main() {
	var err error
	ctx := context.Background()
	// connect
	opts := options.Client().ApplyURI(conn)
	client, err = mongo.Connect(ctx, opts)
	if err != nil {
		panic(err)
	}
	defer client.Disconnect(ctx)

	if err = preflight(ctx); err != nil {
		panic(err)
	}
	// insert
	if err = insert(ctx); err != nil {
		panic(err)
	}
	// query
	if err = query(ctx); err != nil {
		panic(err)
	}
	// update
	if err = update(ctx); err != nil {
		panic(err)
	}
	delete(ctx)
}

func preflight(ctx context.Context) error {
	fmt.Printf("pinging %T ...\n", client)
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		panic(err)
	}
	fmt.Printf("OK\n")
	names, err := client.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		//panic(err)
		return err
	}
	fmt.Println(names)

	db := client.Database("test")
	fmt.Printf("%T\n", db)

	coll = db.Collection("example")
	//defer coll.Drop(ctx)

	fmt.Printf("%T\n", coll)
	return nil
}

func insert(ctx context.Context) error {
	doc := bson.D{
		{"someString", "Example String"},
		{"someInteger", 12},
		{"someStringSlice", []string{"Example 1", "Example 2", "Example 3"}},
	}
	r, err := coll.InsertOne(ctx, doc)
	if err != nil {
		return err
	}
	id = r.InsertedID.(primitive.ObjectID)
	fmt.Println(id)
	docs := []interface{}{
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
	rs, err := coll.InsertMany(ctx, docs)
	if err != nil {
		return err
	}
	fmt.Println(rs.InsertedIDs)
	//time.Sleep(2 * time.Second)
	return nil
}

func query(ctx context.Context) error {
	// M is unordered map
	c := coll.FindOne(ctx, bson.M{"_id": id})
	var result bson.M
	c.Decode(&result)

	fmt.Printf("\nItem with ID: %v contains the following:\n", result["_id"])
	fmt.Println("someString:", result["someString"])
	fmt.Println("someInteger:", result["someInteger"])
	fmt.Println("someStringSlice:", result["someStringSlice"])
	// D is ordered dictionary
	filter := bson.D{{"someInteger", bson.D{{"$lt", 60}}}}
	cursor, err := coll.Find(ctx, filter)
	if err != nil {
		return err
	}
	var docs []bson.M
	if err = cursor.All(ctx, &docs); err != nil {
		return err
	}

	for _, e := range docs {
		fmt.Printf("\nItem with ID: %v contains the following:\n", e["_id"])
		fmt.Println("someString:", e["someString"])
		fmt.Println("someInteger:", e["someInteger"])
		fmt.Println("someStringSlice:", e["someStringSlice"])
	}
	return nil
}

func update(ctx context.Context) error {
	rUpdt, err := coll.UpdateByID(
		ctx,
		id,
		bson.D{
			{"$set", bson.M{"someInteger": 201}},
		},
	)
	if err != nil {
		return err
	}
	fmt.Println("Number of items updated:", rUpdt.ModifiedCount)

	rUpdt, err = coll.UpdateOne(
		ctx,
		bson.M{"_id": id},
		bson.D{
			{"$set", bson.M{"someString": "The Updated String"}},
		},
	)
	if err != nil {
		return err
	}
	fmt.Println("Number of items updated:", rUpdt.ModifiedCount)

	result, err := coll.UpdateMany(
		ctx,
		bson.D{{"someInteger", bson.D{{"$gt", 60}}}},
		bson.D{
			{"$set", bson.M{"someInteger": 60}},
		},
	)
	if err != nil {
		return err
	}
	fmt.Println("Number of items updated:", result.ModifiedCount)
	return nil
}

func delete(ctx context.Context) error {
	result, err := coll.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		return err
	}
	fmt.Println("Number of items deleted:", result.DeletedCount)
	return nil
}
