package beginner

import (
	"fmt"
)

func Loop() {
	var n int 
	fmt.Print("Enter a positive integer N: ")
	if _, err := fmt.Scan(&n); err != nil || n <= 0 {
		fmt.Println("Invalid input. Please enter a positive integer.")
		return 
	}

	sum := 0 
	for i := 0; i <= n; i++ {
		sum += 1
	}

	fmt.Printf("Sum from 1 to %d is: %d\n", n, sum)
}