package aboutSelect

import (
	"fmt"
	"math/rand"
	"time"
)

func MySelect() {
	c := make(chan int, 1)
	go func() {
		r := rand.Intn(300)
		time.Sleep(time.Duration(r) * time.Millisecond)
		c <- r
	}()
	var reslist int
	select {
	case reslist = <-c:
		fmt.Printf("get num is %d\n", reslist)
	case <-time.After(100 * time.Millisecond):
		fmt.Printf("FAIL to get num\n")
	}
}
