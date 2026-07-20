package main
import (
    "bufio"
    "context"
    "fmt"
    "os"
    "strconv"
    "time"
)

func sumWithCtx(ctx context.Context, nums []int) (int, error) {
    total := 0
    for _, n := range nums {
        select {
        case <-ctx.Done(): return 0, ctx.Err() //return immediately if cancelled
        default: //continue
        }
        time.Sleep(time.Microsecond)
        total += n
    }
    return total, nil
}

func main() {
    sc := bufio.NewScanner(os.Stdin)
    sc.Scan(); n, _ := strconv.Atoi(sc.Text())
    nums := make([]int, n)
    for i := 0; i < n; i++ {
        sc.Scan()
        nums[i], _ = strconv.Atoi(sc.Text())
    }
    ctx, cancel := context.WithTimeout(context.Background(), time.Second)
    defer cancel()
    total, err := sumWithCtx(ctx, nums)
    if err != nil {
        fmt.Println("error:", err)
    } else {
        fmt.Println(total)
    }
}
