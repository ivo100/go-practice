package main

import (
	"fmt"
	"log"
	"sort"
)

type Employee struct {
	Name string
	ID   string
	SSN  int
	Age  int
}

func (employee Employee) ToString() string {
	return fmt.Sprintf("%s: %d,%s,%d\n", employee.Name, employee.Age, employee.ID, employee.SSN)
}

// array of people that can be ordered by age
type AgeOrder []Employee

func (sortable AgeOrder) Len() int               { return len(sortable) }
func (sortable AgeOrder) Swap(i int, j int)      { sortable[i], sortable[j] = sortable[j], sortable[i] }
func (sortable AgeOrder) Less(i int, j int) bool { return sortable[i].Age < sortable[j].Age }

func main() {
	var employees = []Employee{
		{"Graham", "100", 111, 31},
		{"John", "200", 222, 42},
		{"Michael", "300", 333, 23},
	}

	log.Printf("\nInital array:\n %+v", employees)

	sort.Sort(AgeOrder(employees))
	log.Printf("\nArray after Sort (ascending order):\n %+v", employees)

	sort.Slice(employees, func(i int, j int) bool {
		return employees[i].Age > employees[j].Age
	})
	log.Printf("\nArray after Slice (descending order):\n %+v", employees)
}
