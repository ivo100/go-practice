package main

import (
	"fmt"
	"github.com/EricFrancis12/stripol"
	"github.com/evandrojr/string-interpolation/esi"
	"github.com/hashicorp/hil"
	"github.com/hashicorp/hil/ast"
	"log"
)

/*
other string interpolation related libraries
https://github.com/avahidi/interpol/tree/master
*/
func test1() {
	input := "${6 + 2}"

	tree, err := hil.Parse(input)
	if err != nil {
		log.Fatal(err)
	}

	result, err := hil.Eval(tree, &hil.EvalConfig{})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Type: %s\n", result.Type)
	fmt.Printf("Value: %s\n", result.Value)

}

func test2() {
	// hasicorp hil can be useful for configuring parameters/external expressions etc.
	input := "${var.test} - ${6 + 2}"
	tree, err := hil.Parse(input)
	if err != nil {
		log.Fatal(err)
	}
	config := &hil.EvalConfig{
		GlobalScope: &ast.BasicScope{
			VarMap: map[string]ast.Variable{
				"var.test": ast.Variable{
					Type:  ast.TypeString,
					Value: "TEST STRING",
				},
			},
		},
	}
	result, err := hil.Eval(tree, config)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Type: %s\n", result.Type)
	fmt.Printf("Value: %s\n", result.Value)
}

func test3() {
	foo := "bar"

	esi.Println("foo ", foo)

	esi.Print("Print ", 10, " ", 7, " interpolates anything ", true, " ", 3.4e10)
	esi.Print(" no line break")
	esi.Println()
	f := esi.Sprint("Sprint ", 10, " ", 7, " interpolates anything ", true, " ", 3.4e10)
	esi.Print(f)
}

func test4() {
	s := stripol.New("${", "}")

	s.RegisterVar("name", "Mike Tyson")
	s.RegisterVar("pet", "tiger")

	str := "${name} has a pet ${pet}."
	result := s.Eval(str)

	fmt.Println(result)
}

func main() {
	//test1()
	//test2()
	//test3()
	test4()
}
