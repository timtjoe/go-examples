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
            result = append(result, x)
        }
    }
    return result
}

func main() {
    sc := bufio.NewScanner(os.Stdin)
    sc.Scan()
    nums := []int{}
    for _, f := range strings.Fields(sc.Text()) {
        n, _ := strconv.Atoi(f)
        nums = append(nums, n)
    }

    evens := Filter(nums, func(n int) bool { return n%2 == 0 })

    parts := make([]string, len(evens))
    for i, v := range evens {
        parts[i] = strconv.Itoa(v)
    }
    fmt.Println(strings.Join(parts, " "))
}
