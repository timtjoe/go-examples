package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Filter[T any](s []T, keep func(T) bool) []T { 
    var result []T 
    for _, x := range s {
        if keep(x) {
            result =append(result, x)
        }
    }
    return result
 }

// Non-generic fallback:
func filterEvens(s []int) []int {
    var r []int
    for _, n := range s {
        if n%2 == 0 {
            r = append(r, n)
        }
    }
    return r
}

func main() {
    sc := bufio.NewScanner(os.Stdin)
    sc.Scan()
    nums := []int{}
    for _, f := range strings.Fields(sc.Text()) {
        n, _ := strconv.Atoi(f)
        nums = append(nums, n)
    }
    evens := filterEvens(nums)
    parts := make([]string, len(evens))
    for i, v := range evens {
        parts[i] = strconv.Itoa(v)
    }
    fmt.Println(strings.Join(parts, " "))
}
