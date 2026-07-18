package intermediate

import (
	"fmt"
	"sync"
)

type SafeCounter struct {
	mu sync.Mutex 
	count int
}

func (c *SafeCounter) Inc(){
	c.mu.Lock()
	defer c.mu.Unlock()
	c.count++
}

func (c *SafeCounter) Get() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.count
}

func Sync() {
	fmt.Println("Hello sync.")
}