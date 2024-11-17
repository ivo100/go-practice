package main

import (
	"bytes"
	"context"
	"fmt"
	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
	"io"
)

func main() {
	//var err error
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		panic(err)
	}
	js, err := jetstream.New(nc)
	if err != nil {
		panic(err)
	}
	ctx := context.Background()

	// Create a new bucket. Bucket name is required and has to be unique within a JetStream account.
	os, err := js.CreateObjectStore(ctx, jetstream.ObjectStoreConfig{Bucket: "configs"})
	if err != nil {
		panic(err)
	}

	config1 := bytes.NewBufferString("first config")
	// Put an object in a bucket. Put expects an object metadata and a reader
	// to read the object data from.
	_, err = os.Put(ctx, jetstream.ObjectMeta{Name: "config-1"}, config1)
	if err != nil {
		panic(err)
	}

	// Objects can also be created using various helper methods

	// 1. As raw strings
	_, err = os.PutString(ctx, "config-2", "second config")
	if err != nil {
		panic(err)
	}

	// 2. As raw bytes
	oi, err := os.PutBytes(ctx, "config-3", []byte("third config"))
	if err != nil {
		panic(err)
	}
	_ = oi

	// 3. As a file
	oi, err = os.PutFile(ctx, "main.go")
	if err != nil {
		panic(err)
	}

	// Get an object
	// Get returns a reader and object info
	// Similar to Put, Get can also be used with helper methods
	// to retrieve object data as a string, bytes or to save it to a file
	object, err := os.Get(ctx, "config-1")
	if err != nil {
		panic(err)
	}
	data, err := io.ReadAll(object)
	if err != nil {
		panic(err)
	}
	info, err := object.Info()
	if err != nil {
		panic(err)
	}

	// Prints `configs.config-1 -> "first config"`
	fmt.Printf("%s.%s -> %q\n", info.Bucket, info.Name, string(data))

	// Delete an object.
	// Delete will remove object data from stream, but object metadata will be kept
	// with a delete marker.
	//os.Delete(ctx, "config-1")

	// getting a deleted object will return an error
	_, err = os.Get(ctx, "config-1")
	fmt.Println(err) // prints `nats: object not found`

	// A bucket can be deleted once it is no longer needed
	//js.DeleteObjectStore(ctx, "configs")
}
