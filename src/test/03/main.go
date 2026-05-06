package main

import "fmt"

type BaseStruct struct {
	Name    string
	ID      int
	Country string
}

type SampleMap map[string]*BaseStruct

func NewSampleMap() SampleMap {
	return SampleMap{
		"first": {
			Name:    "John Doe",
			ID:      1,
			Country: "USA",
		},
		"second": {
			Name:    "Jane Smith",
			ID:      2,
			Country: "UK",
		},
	}
}

func main() {
	m := NewSampleMap()
	for k, v := range m {
		fmt.Printf("%s: %+v\n", k, *v)
	}
}
