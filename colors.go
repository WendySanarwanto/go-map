package main

import "fmt"

type colors map[string]string

func (c colors) print() {
	// TODO: implement this
	fmt.Println("Printed 'colors' map: ")
	for color, hex := range c {
		fmt.Println("\t", color, "color's hex is", hex)
	}
	fmt.Println("--- \n")
}

func (c colors) printValue(key string) {
	fmt.Println(key, " color's hex is", c[key], "\n---\n")
}