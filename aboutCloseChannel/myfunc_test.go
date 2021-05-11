package aboutCloseChannel

import (
	"sync"
	"testing"
)

func TestCloseChannel(t *testing.T) {
	var wg sync.WaitGroup
	ch := make(chan int, 1)
	wg.Add(1)
	dataProducer(ch, &wg)
	wg.Add(1)
	dataReciver(ch, &wg)
	wg.Wait()
}
