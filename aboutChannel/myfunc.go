package aboutChannel

import (
	"log"
	"sync"
)

func Producer(ch chan int, wg *sync.WaitGroup) {
	go func() {
		for i := 0; i < 100; i++ {
			ch <- i
		}
		close(ch)
		wg.Done()
	}()
}
func Consumer(ch chan int, wg *sync.WaitGroup) {
	go func() {
		for {
			if data, ok := <-ch; ok {
				log.Printf("%d\n", data)
			} else {
				break
			}
		}
		wg.Done()
	}()
}
