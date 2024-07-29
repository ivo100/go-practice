package wheel

type Wheel struct {
	Name string
}

func NewWheel(name string) *Wheel {
	return &Wheel{Name: name}
}

func (w *Wheel) String() string {
	return "Wheel: " + w.Name
}
