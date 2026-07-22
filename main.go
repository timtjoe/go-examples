package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
)

var ErrTooSmall = errors.New("value too small")

func validate(n int) error { 
    if n >= 10 {
        return nil
    }
    return fmt.Errorf("validating n=%d: %w", n, ErrTooSmall)
}

func main() {
    sc := bufio.NewScanner(os.Stdin)
    sc.Scan()
    n, _ := strconv.Atoi(sc.Text())

    err := validate(n)
    if err == nil {
        fmt.Println("ok")
    } else if errors.Is(err, ErrTooSmall) {
        fmt.Printf("too small: %d\n", n)
    } else {
    fmt.Printf("error: %v\n", err)
    }
}
