package beginner

import "fmt"

func Slices() {
	// base := []int{}
	nums := []int{3, 8, 5, 9, 1}
	max := nums[0]
	for _, value := range nums[1:] {
		fmt.Printf("inner: %d\n", value)
		if value > max {
			max = value
		}
	}

	fmt.Printf("outer: %d", max)
}
