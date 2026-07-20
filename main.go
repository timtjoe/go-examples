package main
import (
    "fmt"
    "sync"
    "sync/atomic"
)

func main() {
    var counter int64
    var wg sync.WaitGroup
    // spawn 100 goroutines
    for i := 0; i < 100; i++ {
        wg.Add(1);
        go func()  {
            defer wg.Done()
            // each does 100 atomic.AddInt64(&counter, 1)
            for j := 0; j < 100; j++ {
                atomic.AddInt64(&counter, 1)
            }
        }()
    }
    wg.Wait()
    
    fmt.Println(counter)
}
