/*
********************************************************************************
Golang - Asterisk and Ampersand Cheatsheet
********************************************************************************
Also available at: https://play.golang.org/p/lNpnS9j1ma
Allowed:
--------
p := Person{"Steve", 28} 	stores the value
p := &Person{"Steve", 28} 	stores the pointer address (reference)
PrintPerson(p) 			passes either the value or pointer address (reference)
PrintPerson(*p) 		passes the value
PrintPerson(&p) 		passes the pointer address (reference)
func PrintPerson(p Person)	ONLY receives the value
func PrintPerson(p *Person)	ONLY receives the pointer address (reference)
Not Allowed:
--------
p := *Person{"Steve", 28} 	illegal
func PrintPerson(p &Person)	illegal
*/

package main

import (
	"fmt"
)

//Person persona
type Person struct {
	Name string
	Age  int
}

// This only works with *Person, does not work with Person
// Only works with Test 2 and Test 3
func (p *Person) String() string {
	return fmt.Sprintf("%s is %d", p.Name, p.Age)
}

// This works with both *Person and Person, BUT you can't modiy the value and
// it takes up more space
// Works with Test 1, Test 2, Test 3, and Test 4
/*func (p Person) String() string {
	return fmt.Sprintf("%s is %d", p.Name, p.Age)
}*/

// *****************************************************************************
// Test 1 - Pass by Value
// *****************************************************************************

func test1() {
	p := Person{"Steve", 28}
	printPerson1(p)  //String: {Steve 28} | Name: Steve | Age: 28
	updatePerson1(p) //String: {Steve 32} | Name: Steve | Age: 32
	printPerson1(p)  //String: {Steve 28} | Name: Steve | Age: 28
}

func updatePerson1(p Person) {
	p.Age = 32
	printPerson1(p)
}

func printPerson1(p Person) {
	fmt.Printf("String: %v | Name: %v | Age: %d\n",
		p,
		p.Name,
		p.Age)
}

// *****************************************************************************
// Test 2 - Pass by Reference
// *****************************************************************************

func test2() {
	p := &Person{"Steve", 28}
	printPerson2(p)  //String: Steve is 28 | Name: Steve | Age: 28
	updatePerson2(p) //String: Steve is 32 | Name: Steve | Age: 32
	printPerson2(p)  //String: Steve is 32 | Name: Steve | Age: 32
}

func updatePerson2(p *Person) {
	p.Age = 32
	printPerson2(p)
}

func printPerson2(p *Person) {
	fmt.Printf("String: %v | Name: %v | Age: %d\n",
		p,
		p.Name,
		p.Age)
}

// *****************************************************************************
// Test 3 - Pass by Reference (requires more typing)
// *****************************************************************************

func test3() {
	p := Person{"Steve", 28}
	printPerson3(&p)  //String: Steve is 28 | Name: Steve | Age: 28
	updatePerson3(&p) //String: Steve is 32 | Name: Steve | Age: 32
	printPerson3(&p)  //String: Steve is 32 | Name: Steve | Age: 32
}

func updatePerson3(p *Person) {
	p.Age = 32
	printPerson3(p)
}

func printPerson3(p *Person) {
	fmt.Printf("String: %v | Name: %v | Age: %d\n",
		p,
		p.Name,
		p.Age)
}

// *****************************************************************************
// Test 4 - Pass by Value (requires more typing)
// *****************************************************************************

func test4() {
	p := &Person{"Steve", 28}
	printPerson4(*p)  //String: {Steve 28} | Name: Steve | Age: 28
	updatePerson4(*p) //String: {Steve 32} | Name: Steve | Age: 32
	printPerson4(*p)  //String: {Steve 28} | Name: Steve | Age: 28
}

func updatePerson4(p Person) {
	p.Age = 32
	printPerson4(p)
}

func printPerson4(p Person) {
	fmt.Printf("String: %v | Name: %v | Age: %d\n",
		p,
		p.Name,
		p.Age)
}

// *****************************************************************************
// Main
// *****************************************************************************

func main() {
	test1()
	test2()
	test3()
	test4()
}
