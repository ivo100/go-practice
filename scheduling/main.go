package main

import (
	"context"
	"time"

	"github.com/adhocore/gronx/pkg/tasker"
)

func main() {
	// run task without overlap, set concurrent flag to false:
	concurrent := true
	taskr := tasker.New(tasker.Option{
		Verbose: true,
		// optional: defaults to local
		//Tz: "Asia/Bangkok",
		// optional: defaults to stderr log stream
		//Out: "/tmp/tasker.log",
	})
	taskr.Task("*/2 * * * *", func(ctx context.Context) (int, error) { // every 2 minutes
		taskr.Log.Printf("task started, working...")
		time.Sleep(180 * time.Second)
		taskr.Log.Printf("task finished")
		return 0, nil
	}, concurrent)

	taskr.Task("@5minutes", taskr.Taskify("ls -al", tasker.Option{}), concurrent)
	// every 10 minute with arbitrary command
	//taskr.Task("@10minutes", taskr.Taskify("command --option val -- args", tasker.Option{Shell: "/bin/sh -c"}))

	// ... add more tasks

	// optionally if you want tasker to stop after 2 hour, pass the duration with Until():
	taskr.Until(10 * time.Minute)

	// finally run the tasker, it ticks sharply on every minute and runs all the tasks due on that time!
	// it exits gracefully when ctrl+c is received making sure pending tasks are completed.
	taskr.Run()
}
