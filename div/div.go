package main

import "fmt"

func main() {
	/*
		fmt.Println(div(7, 3))
		fmt.Println(div(7, 0)) // panic
	*/
	fmt.Println(safeDiv(7, 3))
	fmt.Println(safeDiv(7, 0))
}

func safeDiv(x, y int) (q int, err error) {
	// q & err are local variables (like x & y)
	// catch/except
	defer func() {
		if e := recover(); e != nil {
			// fmt.Printf("error: %#v\n", e)
			err = fmt.Errorf("%v", e)
		}
	}()

	// q = div(x, y)
	// return
	return div(x, y), nil
}

// div is from external package
func div(x, y int) int {
	return x / y
}
