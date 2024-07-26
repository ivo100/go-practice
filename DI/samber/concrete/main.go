package main

import (
	"fmt"
	"github.com/samber/do/v2"
	"os"
	"syscall"
)

type Wheel struct {
}

type Engine struct {
	Started bool
}

func (e *Engine) Shutdown() error {
	// called on injector shutdown
	e.Started = false
	return nil
}

type Car struct {
	Engine *Engine
	Wheels []*Wheel
}

func (c *Car) Start() {
	c.Engine.Started = true
	println("vroooom")
}

func main() {
	fmt.Println("Start DI")

	di := do.New()

	// provide wheels
	do.ProvideNamedValue(di, "wheel-1", &Wheel{})
	do.ProvideNamedValue(di, "wheel-2", &Wheel{})
	do.ProvideNamedValue(di, "wheel-3", &Wheel{})
	do.ProvideNamedValue(di, "wheel-4", &Wheel{})

	// provide car
	do.Provide(di, func(i do.Injector) (*Car, error) {
		car := Car{
			Engine: do.MustInvoke[*Engine](i),
			Wheels: []*Wheel{
				do.MustInvokeNamed[*Wheel](i, "wheel-1"),
				do.MustInvokeNamed[*Wheel](i, "wheel-2"),
				do.MustInvokeNamed[*Wheel](i, "wheel-3"),
				do.MustInvokeNamed[*Wheel](i, "wheel-4"),
			},
		}
		return &car, nil
	})

	// provide engine
	do.Provide(di, func(i do.Injector) (*Engine, error) {
		return &Engine{}, nil
	})

	// start the car
	car := do.MustInvoke[*Car](di)
	car.Start()

	// will block - handle ctrl-c and shutdown services
	di.ShutdownOnSignals(syscall.SIGTERM, os.Interrupt)
}
