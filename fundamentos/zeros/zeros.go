package main

import "fmt"

func main() {
	fmt.Println("Olá mundo!")

	var a int
	var b float64
	var c bool
	var d string
	var e *int

	fmt.Printf("%v %v %v %q %v\n", a, b, c, d, e)
}
