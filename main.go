package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "sync"
)

func main() {
    sc := bufio.NewScanner(os.Stdin)
    sc.Buffer(make([]byte, 1024*1024), 1024*1024)
    sc.Scan()
    n, _ := strconv.Atoi(sc.Text())
    nums := make([]int, n)
    for i := 0; i < n; i++ {
        sc.Scan()
        nums[i], _ = strconv.Atoi(sc.Text())
    }

    var mu sync.Mutex
    var wg sync.WaitGroup
    total := 0

    // Split into 4 chunks
    chunkSize := (n + 3) / 4 // ceiling division
    for start := 0; start < n; start += chunkSize {
        end := start + chunkSize
        if end > n {
            end = n
        }
        wg.Add(1)
        go func(slice []int) {
            defer wg.Done()
            sum := 0
            for _, v := range slice {
                sum += v
            }
            mu.Lock()
            total += sum
            mu.Unlock()
        }(nums[start:end])
    }

    // Ensure exactly 4 goroutines even if n < 4
    for g := len(nums)/chunkSize + 1; g < 4; g++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            // empty slice, does nothing
        }()
    }

    wg.Wait()
    fmt.Println(total)
}
