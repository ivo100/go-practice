package engine

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
