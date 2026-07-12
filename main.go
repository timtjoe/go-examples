package main

import (
	"fmt"
	beginner "go-examples/beginner"
	"log"
)

func main() {

	// Named return values
	nums := []int{1,2,3,4}
	min, max := beginner.MinMax(nums)
	fmt.Printf("Min: %d Max: %d \n", min, max)

	// The error-pair idiom
	age, err := beginner.ParseAge("42")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(age)

	// Discarding returns with _ 
	beginner.Discarding()
}
