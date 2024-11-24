package main

import (
	"context"
	"fmt"
	"github.com/evandrojr/string-interpolation/esi"
	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
	"time"
)

/*
JetStream KeyValue Stores offer a straightforward method for storing key-value pairs within JetStream.
These stores are supported by a specially configured stream, designed to efficiently and compactly store these pairs.
This structure ensures rapid and convenient access to the data.

The KV Store, also known as a bucket, enables the execution of various operations:

	create/update a value for a given key
	get a value for a given key
	delete a value for a given key
	purge all values from a bucket
	list all keys in a bucket
	watch for changes on given key set or the whole bucket
	retrieve history of changes for a given key
*/
func main() {
	simple()
	//watch()
}

func simple() {
	nc, _ := nats.Connect(nats.DefaultURL)
	js, _ := jetstream.New(nc)
	// In the `jetstream` package, almost all API calls rely on `context.Context` for timeout/cancellation handling
	ctx := context.Background()
	// Create a new KV store. Bucket name is required and has to be unique within a JetStream account.
	// idempotent as long as config is the same
	kv, _ := js.CreateKeyValue(ctx, jetstream.KeyValueConfig{
		Bucket: "testbucket1",
		TTL:    1 * time.Minute,
	})
	// Set a value for a given key
	// Put will either create or update a value for a given key
	kv.Put(ctx, "sue.color", []byte("blue"))

	// Get an entry for a given key
	// Entry contains key/value, but also metadata (revision, timestamp, etc.))
	entry, err := kv.Get(ctx, "sue.color")
	if err != nil {
		fmt.Println(err)
	}
	esi.Println(entry.Key(), " ", entry.Revision(), " ", string(entry.Value()))
	kv.Put(ctx, "sue.color", []byte("red"))
	entry, err = kv.Get(ctx, "sue.color")
	if err != nil {
		fmt.Println(err)
	}

	// Prints `sue.color @ 1 -> "blue"`
	fmt.Printf("%s @ %d -> %q\n", entry.Key(), entry.Revision(), string(entry.Value()))

	// Update a value for a given key
	// Update will fail if the key does not exist or the revision has changed
	kv.Update(ctx, "sue.color", []byte("red"), 1)

	// Create will fail if the key already exists
	_, err = kv.Create(ctx, "sue.color2", []byte("purple"))
	fmt.Println(err) // prints `nats: key exists`

	keys, _ := kv.ListKeys(ctx)

	// Prints all 3 keys
	for key := range keys.Keys() {
		fmt.Println(key)
	}

	// Purge will remove all keys from a bucket.
	// The latest revision of each key will be kept
	// with a delete marker, all previous revisions will be removed
	// permanently.
	//for key := range keys.Keys() {
	//	kv.Purge(ctx, key)
	//}
	// PurgeDeletes will remove all keys from a bucket
	// with a delete marker.
	kv.PurgeDeletes(ctx)
	// Delete a value for a given key.
	// Delete is not destructive, it will add a delete marker for a given key
	// and all previous revisions will still be available
	kv.Delete(ctx, "sue.color")

	// getting a deleted key will return an error
	_, err = kv.Get(ctx, "sue.color")
	fmt.Println(err) // prints `nats: key not found`

	// A bucket can be deleted once it is no longer needed
	//js.DeleteKeyValue(ctx, "profiles")

}

func watch() {
	nc, _ := nats.Connect(nats.DefaultURL)

	js, _ := jetstream.New(nc)
	ctx := context.Background()
	kv, _ := js.CreateKeyValue(ctx, jetstream.KeyValueConfig{
		Bucket: "profiles",
		TTL:    24 * time.Hour,
	})

	kv.Put(ctx, "sue.color", []byte("blue"))

	// A watcher can be created to watch for changes on a given key or the whole bucket
	// By default, watcher will return most recent values for all matching keys.
	// Watcher can be configured to only return updates by using jetstream.UpdatesOnly() option.
	watcher, _ := kv.Watch(ctx, "sue.*")
	defer watcher.Stop()

	kv.Put(ctx, "sue.age", []byte("43"))
	kv.Put(ctx, "sue.color", []byte("red"))

	// First, the watcher sends most recent values for all matching keys.
	// In this case, it will send a single entry for `sue.color`.
	entry := <-watcher.Updates()
	// Prints `sue.color @ 1 -> "blue"`
	fmt.Printf("%s @ %d -> %q\n", entry.Key(), entry.Revision(), string(entry.Value()))

	// After all current values have been sent, watcher will send nil on the channel.
	entry = <-watcher.Updates()
	if entry != nil {
		fmt.Println("Unexpected entry received")
	}

	// After that, watcher will send updates when changes occur
	// In this case, it will send an entry for `sue.color` and `sue.age`.

	entry = <-watcher.Updates()
	// Prints `sue.age @ 2 -> "43"`
	fmt.Printf("%s @ %d -> %q\n", entry.Key(), entry.Revision(), string(entry.Value()))

	entry = <-watcher.Updates()
	// Prints `sue.color @ 3 -> "red"`
	fmt.Printf("%s @ %d -> %q\n", entry.Key(), entry.Revision(), string(entry.Value()))

	kv.Put(ctx, "sue.color", []byte("blue"))
	kv.Put(ctx, "sue.age", []byte("43"))
	kv.Put(ctx, "bucket", []byte("profiles"))

}
