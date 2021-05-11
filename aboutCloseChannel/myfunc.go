package aboutCloseChannel

import (
	"fmt"
	"sync"
)

func dataProducer(ch chan int, wg *sync.WaitGroup) {
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		close(ch)
		wg.Done()
	}()
}
func dataReciver(ch chan int, wg *sync.WaitGroup) {
	go func() {
		for {
			if data, ok := <-ch; ok {
				fmt.Printf("%d\n", data)
			} else {
				break
			}

		}
		wg.Done()
	}()
}
