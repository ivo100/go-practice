package car

import "di/engine"

type Car struct {
	Engine *engine.Engine
	Wheels []*engine.Wheel
}

func (c *Car) Start() {
	c.Engine.Started = true
	println("vroooom")
}
