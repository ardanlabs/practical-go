package main

import "fmt"

func main() {
	// var i interface{}
	var i any // go >= 1.18

	i = 7
	fmt.Println(i)

	i = "hello"
	fmt.Println(i)

	s := i.(string) // type assertion
	fmt.Println("s:", s)

	/*
		n := i.(int) // will panic
		fmt.Println("n:", n)
	*/
	n, ok := i.(int) // comma ok
	fmt.Println("n:", n, "ok:", ok)
}
