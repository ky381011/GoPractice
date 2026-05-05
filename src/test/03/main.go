package main

import "fmt"

type baseStruct struct {
	Name    string
	ID      int
	Country string
}

type SampleMap map[string]baseStruct

func NewSampleMap() *SampleMap {
	return &SampleMap{
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
	fmt.Println(NewSampleMap())
}
