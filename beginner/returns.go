package beginner

import (
	"fmt"
	"os"
	"strconv"
)

func divmod(a, b int) (int, int) {
	return a / b, a % b
}

// Retunring multiple values
func Return_Multi() {
	q, r := divmod(17, 5)
	fmt.Println(q, r) //3 2
}

// Named return values - for self-documenting signatures
func MinMax(nums []int)(min, max int) {
	min = nums[0]
	max = nums[0]
	for _, n := range nums[1:] {
		if n < min {min = n}
		if n > max {max = n}
	}
	return // "naked return" -- uses the name values. same as return min, max
}

// The error-pair idiom - Idiomatic Go's error handling
func ParseAge(s string) (int, error) {
	n, err := strconv.Atoi(s)
	if err != nil {
		return 0, fmt.Errorf("Invalid age %q: %w", s, err)
	}
	if n < 0 {
		return 0, fmt.Errorf("Age cannot be negative: %d", n)
	}
	return n, nil
} 

// Discarding return with _ - when you don't need a value, use _:
func Discarding() {
	// Don't care about the error:
	hostname, _ := os.Hostname()
	println("Hostname: ", hostname)

	// Don't care about the quotient, just remainder:
	_, r := divmod(17, 5)
	fmt.Println(r)
}

// Panic and recover (briefly) 
// Panic for truly unrecoverable situations (programmer errors), Go has panic:
func MustParse(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return n
}
