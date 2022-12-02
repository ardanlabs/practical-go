package main

import "fmt"

func main() {
	fmt.Println(Max([]int{1, 2, 3}))
	fmt.Println(Max([]float64{1, 2, 3}))
}

type Number interface {
	int | float64
}

func Max[T Number](values []T) (T, error) {
	if len(values) == 0 {
		return 0, fmt.Errorf("MaxFloat64s of empty slice")
	}

	max := values[0]
	for _, val := range values[1:] {
		if val > max {
			max = val
		}
	}
	return max, nil
}

/*
func MaxInts(values []int) (int, error) {
	if len(values) == 0 {
		return 0, fmt.Errorf("MaxInts of empty slice")
	}

	max := values[0]
	for _, val := range values[1:] {
		if val > max {
			max = val
		}
	}
	return max, nil
}

func MaxFloat64s(values []float64) (float64, error) {
	if len(values) == 0 {
		return 0, fmt.Errorf("MaxFloat64s of empty slice")
	}

	max := values[0]
	for _, val := range values[1:] {
		if val > max {
			max = val
		}
	}
	return max, nil
}

*/
