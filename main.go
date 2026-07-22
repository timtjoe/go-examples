package main

import "fmt"

func add(a, b int) int { 
    return a + b;
}

func main() {
    cases := []struct {
        name string
        a, b int
        want int
    }{
        {"two positives", 2, 3, 5},
        {"with zero", 0, 7, 7},
        {"with negative", -2, -3, -5},
    }
    for _, c := range cases {
        got := add(c.a, c.b)
        if got == c.want { 
            fmt.Printf("%s: PASS\n", c.name)
        } else { 
            fmt.Printf("%s: FAIL got=%d want=%d\n", c.name, got, c.want)
        }
    }
}
