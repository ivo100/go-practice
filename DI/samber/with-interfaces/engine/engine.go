package engine

import (
	"errors"
	"fmt"
	"github.com/samber/do/v2"
)

type Engine interface {
	Start() error
	Shutdown() error
}

type EngImpl struct {
	Started bool
}

func NewEngine(i do.Injector) (*EngImpl, error) {
	return &EngImpl{}, nil
}

func (e *EngImpl) Start() error {
	if e.Started {
		return errors.New("already started")
	}
	fmt.Println("engine starts...")
	e.Started = true
	return nil
}

// Shutdown is called on injector shutdown
func (e *EngImpl) Shutdown() error {
	if !e.Started {
		return errors.New("not started")
	}
	fmt.Println("Shutdown")
	e.Started = false
	return nil
}
