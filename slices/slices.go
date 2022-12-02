package main

import (
	"fmt"
	"sort"
)

func main() {
	s1 := make([]int, 10)
	fmt.Printf("s1: len=%d, cap=%d\n", len(s1), cap(s1))
	s2 := s1[3:7]
	fmt.Printf("s2: len=%d, cap=%d\n", len(s2), cap(s2))
	s2 = append(s2, 7)
	fmt.Println("s1:", s1)
	fmt.Println("s2:", s2)
	fmt.Printf("s2: len=%d, cap=%d\n", len(s2), cap(s2))

	arr := [3]int{1, 2, 3}
	fmt.Printf("arr: %v (%T)\n", arr, arr)

	var s []int
	for i := 0; i < 100; i++ {
		s = appendInt(s, i)
	}
	fmt.Println(s[:10])

	fmt.Println(concat([]string{"A", "B"}, []string{"C"}))

	values := []float64{3, 1, 2}
	fmt.Println(median(values)) // 2
	values = []float64{3, 1, 2, 4}
	fmt.Println(median(values)) // 2.5

	fmt.Println(values)
}

/*
median:
- sort values
- if odd number of values, return middle
- otherwise return average of middles

Hint: sort.Float64s
*/
func median(values []float64) float64 {
	// Don't change input slice
	vals := make([]float64, len(values))
	copy(vals, values)

	sort.Float64s(vals)
	i := len(vals) / 2
	if len(vals)%2 == 1 {
		return vals[i]
	}

	return (vals[i-1] + vals[i]) / 2
}

func concat(s1, s2 []string) []string {
	// return append(s1, s2...)
	s := make([]string, len(s1)+len(s2))
	copy(s, s1)
	copy(s[len(s1):], s2)
	return s
}

func appendInt(values []int, val int) []int {
	i := len(values) // where to insert val
	if len(values) < cap(values) {
		values = values[:len(values)+1]
	} else { // re-allocate & copy
		size := len(values)*2 + 1
		fmt.Printf("%d -> %d\n", cap(values), size)
		s1 := make([]int, size)
		copy(s1, values)
		values = s1[:len(values)+1]
	}

	values[i] = val
	return values
}
