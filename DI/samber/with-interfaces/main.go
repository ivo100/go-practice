package main

import (
	"di/with-interfaces/car"
	"di/with-interfaces/engine"
	"di/with-interfaces/wheel"
	"fmt"
	"github.com/samber/do/v2"
	"os"
	"syscall"
)

/*
USING INTERFACES instead of struct
*/

func mainI() {
	fmt.Println("Start DI")

	di := do.New()

	// provide wheels
	do.ProvideNamedValue(di, "wheel-1", wheel.NewWheel("w1"))
	do.ProvideNamedValue(di, "wheel-2", wheel.NewWheel("w2"))
	do.ProvideNamedValue(di, "wheel-3", wheel.NewWheel("w3"))
	do.ProvideNamedValue(di, "wheel-4", wheel.NewWheel("w4"))

	// provide car
	do.Provide(di, car.NewCar)

	// provide engine
	do.Provide(di, engine.NewEngine)

	// DI magic happens here
	do.As[*car.CarImpl, car.Car](di)
	do.As[*engine.EngImpl, engine.Engine](di)
	fmt.Println("root scope -->", di.ID(), di.ListProvidedServices())

	// start the car
	ford := do.MustInvoke[car.Car](di)
	ford.Start()

	// should error
	err := ford.Start()
	if err != nil {
		fmt.Println(err)
	}

	//prius := do.MustInvoke[car.Car](di)
	//err = prius.Start()
	//if err != nil {
	//	fmt.Println(err)
	//} else {
	//	time.Sleep(2 * time.Second)
	//	prius.Stop()
	//}

	// will block - handle ctrl-c and shutdown services
	di.ShutdownOnSignals(syscall.SIGTERM, os.Interrupt)
}
