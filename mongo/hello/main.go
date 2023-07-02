package main

import (
	"context"
	"fmt"
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

}
