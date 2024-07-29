package main

import (
	"fmt"
	"github.com/samber/do/v2"
)

type Foo struct {
	Name string
}

func (f *Foo) String() string {
	return fmt.Sprintf("Foo: %s ", f.Name)
}

func NewFoo(name string) *Foo {
	return &Foo{Name: name}
}

type Bar struct {
	Name string
	Foo  *Foo
}

type Bar2 struct {
	Name string
	Foo  *Foo `do:""` // <- injected by MustInvokeStruct
}

func NewBar(name string) *Bar {
	return &Bar{Name: name}
}

func (b *Bar) WithFoo(foo *Foo) *Bar {
	b.Foo = foo
	return b
}

func (b *Bar) String() string {
	s := fmt.Sprintf("Bar: %s ", b.Name)
	if b.Foo != nil {
		s += fmt.Sprintf(", %s", b.Foo.String())
	}
	return s
}

func (b *Bar2) String() string {
	s := fmt.Sprintf("Bar2: %s ", b.Name)
	if b.Foo != nil {
		s += fmt.Sprintf(", %s", b.Foo.String())
	}
	return s
}

func main() {
	di := do.New()

	do.ProvideNamedValue(di, "f1", NewFoo("a"))
	//do.ProvideValue(di, NewFoo("a"))

	do.ProvideNamed(di, "b1", func(i do.Injector) (*Bar, error) {
		bar := Bar{
			Name: "Bar1",
			//Foo: do.MustInvoke[*Foo](i),
			Foo: do.MustInvokeNamed[*Foo](i, "f1"),
		}
		return &bar, nil
	})

	bar1 := do.MustInvokeNamed[*Bar](di, "b1")
	fmt.Println(bar1)

	// when using non-named - default name is *main.Foo, *main.Bar - one inst per package
	do.ProvideValue(di, NewFoo("b"))
	do.Provide(di, func(i do.Injector) (*Bar, error) {
		bar := Bar{
			Name: "Bar",
			Foo:  do.MustInvoke[*Foo](i),
		}
		return &bar, nil
	})

	bar := do.MustInvoke[*Bar](di)
	fmt.Println(bar)

	do.Provide(di, func(i do.Injector) (*Bar2, error) {
		b := Bar2{
			Name: "Bar2",
			Foo:  do.MustInvokeStruct[Foo](i), // no ctor used, just struct
		}
		b.Foo.Name = "baz"
		return &b, nil
	})

	bar2 := do.MustInvoke[*Bar2](di)
	fmt.Println(bar2.String())
}
