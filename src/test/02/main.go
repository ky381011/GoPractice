package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

var person = Person{
	Name: "Alice",
	Age:  30,
}

func main() {
	fmt.Println(person)
}
