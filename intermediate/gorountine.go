/*
A gorountine is a function called with go keyword and schedule by go scheduler for execution on the preallocated OS thread.
*/

package intermediate

import (
	"fmt"
	"sync"
)

func Say(msg string) {
	fmt.Println(msg)
}

// Cordinating with sync.WaitGroup:

var wg sync.WaitGroup 

func SyncWait(){
	for i := range 100 {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			fmt.Println("Worker", n)
		}(i)
	}
	wg.Wait()
}

// Sharing data: mutexes 
var mu sync.Mutex
var counter int 

func SyncMutex(){
	go func() {
		mu.Lock()
		counter++
		mu.Unlock()
	}()
}