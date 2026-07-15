// Channels are typed pipes for gorountines to talk:
package intermediate

import "fmt"

// Producer/Consumer pattern:
func producer(ch chan<-int){ //chan<- = send-only
	defer close(ch)
	for i := range 5 {
		ch <- i * i
	}
}

func RunChan(){
	ch := make(chan int)
	go producer(ch)
	for v := range ch {
		fmt.Println(v)
	}

// Select - waits on multiple channels:4
// 	select {
// 	case v := <-ch1: fmt.Println("from ch1: ", v)
// 	}
// }
