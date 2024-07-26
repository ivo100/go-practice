package main

import (
	"di/car"
	"di/engine"
	"fmt"
	"github.com/samber/do/v2"
	"os"
	"syscall"
)

func main() {
	fmt.Println("Start DI")

	di := do.New()

	// provide wheels
	do.ProvideNamedValue(di, "wheel-1", &engine.Wheel{})
	do.ProvideNamedValue(di, "wheel-2", &engine.Wheel{})
	do.ProvideNamedValue(di, "wheel-3", &engine.Wheel{})
	do.ProvideNamedValue(di, "wheel-4", &engine.Wheel{})

	// provide car
	do.Provide(di, func(i do.Injector) (*car.Car, error) {
		car := car.Car{
			Engine: do.MustInvoke[*engine.Engine](i),
			Wheels: []*engine.Wheel{
				do.MustInvokeNamed[*engine.Wheel](i, "wheel-1"),
				do.MustInvokeNamed[*engine.Wheel](i, "wheel-2"),
				do.MustInvokeNamed[*engine.Wheel](i, "wheel-3"),
				do.MustInvokeNamed[*engine.Wheel](i, "wheel-4"),
			},
		}
		return &car, nil
	})

	// provide engine
	do.Provide(di, func(i do.Injector) (*engine.Engine, error) {
		return &engine.Engine{}, nil
	})

	// start the car
	car := do.MustInvoke[*car.Car](di)
	car.Start()

	// will block - handle ctrl-c and shutdown services
	di.ShutdownOnSignals(syscall.SIGTERM, os.Interrupt)
}
