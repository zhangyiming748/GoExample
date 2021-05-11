package aboutChannel

import (
	"log"
	"sync"
	"testing"
)

func TestMyChannle(t *testing.T) {
	var wg sync.WaitGroup
	ch := make(chan int, 1)
	wg.Add(1)
	Producer(ch, &wg)
	wg.Add(1)
	Consumer(ch, &wg)
	wg.Wait()

}
func TestAllInOne(t *testing.T) {
	ch := make(chan int, 1)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		close(ch)
		wg.Done()

	}()
	wg.Add(1)
	go func() {
		for {
			if data, ok := <-ch; ok {
				log.Printf("data is %d\n", data)
			} else {
				break
			}
		}
		wg.Done()
	}()
	wg.Wait()
}
