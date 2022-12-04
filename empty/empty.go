package main

import "fmt"

func main() {
	var i any
	// go < 1.18
	// var i interface{}

	i = 7
	fmt.Println(i)

	i = "hi"
	fmt.Println(i)

	// Rule of thumb: Don't use any :)

	s := i.(string) // type assertion
	fmt.Println("s:", s)

	/*
		n := i.(int) // will panic
		fmt.Println("n:", n)
	*/

	// comma, ok
	n, ok := i.(int)
	if ok {
		fmt.Println("n:", n)
	} else {
		fmt.Println("not an int")
	}

	switch i.(type) { // type switch
	case int:
		fmt.Println("an int")
	case string:
		fmt.Println("a string")
	default:
		fmt.Printf("unknown type: %T\n", i)
	}

	/*
		fmt.Println(maxInts([]int{3, 1, 2}))
		fmt.Println(maxFloat64s([]float64{3, 1, 2}))
	*/
	fmt.Println(max([]int{3, 1, 2}))
	fmt.Println(max([]float64{3, 1, 2}))
}

type Number interface {
	int | float64
}

// func max[T int | float64](nums []T) T {
func max[T Number](nums []T) T {
	if len(nums) == 0 {
		return 0
	}

	max := nums[0]
	for _, n := range nums[1:] {
		if n > max {
			max = n
		}
	}
	return max
}

func maxInts(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	max := nums[0]
	for _, n := range nums[1:] {
		if n > max {
			max = n
		}
	}
	return max
}

func maxFloat64s(nums []float64) float64 {
	if len(nums) == 0 {
		return 0
	}

	max := nums[0]
	for _, n := range nums[1:] {
		if n > max {
			max = n
		}
	}
	return max
}
