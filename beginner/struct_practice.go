package beginner

/*
Point Distance
Define a struct Point with X and Y as ints. Read four integers from stdin (one per line) — x1, y1, x2, y2. Build two Points and print the squared Euclidean distance: (x2-x1)² + (y2-y1)².
*/

import (
	"fmt"
)

// Define a struct Point with X and Y as ints.
type Point struct {
	X, Y int
}

func squared(p1, p2 Point) int {
	dx := p2.X - p1.X
	dy := p1.Y - p2.Y
	return dx*dx + dy*dy
}

func StructPractice(){
	var x1, x2, y1, y2 int
	fmt.Print("Enter x1 y1 x2 y2: ")

	// read four integer from  stdin
	fmt.Scan(&x1, &y1, &x2, &y2)

	// Build two pointers
	p1 := Point{X: x1, Y: y1}
	p2 := Point{X: x2, Y: y2}
	
	// square and print euclidean distance
	result := squared(p1, p2)
	fmt.Printf("Squared euclidean distance: %d\n", result);
}

