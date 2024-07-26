package car

import (
	"di/with-interfaces/engine"
	"di/with-interfaces/wheel"
	"fmt"
	"github.com/samber/do/v2"
)

type Car interface {
	Start() error
	Stop() error
}

type CarImpl struct {
	Engine engine.Engine
	Wheels []*wheel.Wheel
}

func NewCar(i do.Injector) (*CarImpl, error) {
	wheels := []*wheel.Wheel{
		do.MustInvokeNamed[*wheel.Wheel](i, "wheel-1"),
		do.MustInvokeNamed[*wheel.Wheel](i, "wheel-2"),
		do.MustInvokeNamed[*wheel.Wheel](i, "wheel-3"),
		do.MustInvokeNamed[*wheel.Wheel](i, "wheel-4"),
	}

	engine := do.MustInvoke[engine.Engine](i)

	car := CarImpl{
		Engine: engine,
		Wheels: wheels,
	}

	return &car, nil
}

func (c CarImpl) Start() error {
	fmt.Println("start the car")
	return c.Engine.Start()
}

func (c CarImpl) Stop() error {
	fmt.Println("stop the car")
	return c.Engine.Shutdown()
}
