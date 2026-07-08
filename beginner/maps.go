// Map are Go's hash tables. Keys must be comparable types; values can be anything

package beginner

import (
	"fmt"
	"strings"
)

func PrintAges() {
	ages := map[string]int{
		"Alice": 30,
		"Bob":   25,
		"Dan":   40,
	}
	if age, ok := ages["Dan"]; ok {
		fmt.Println("ok", ok)
		fmt.Println("found", age)
	}
}

func PrintDistinct() {
	// space-separated words
	line := "apple banana apple orange banana"
	// splite into words
	words := strings.Fields(line)
	// Add each word and print current size
	seen := make(map[string]bool)
	for _, word := range words {
		seen[word] = true
		fmt.Printf("Added: %-7s | Distinct: %d\n", word, len(seen))
	}

}
