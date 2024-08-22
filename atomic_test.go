package belajar_golang_goroutine

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
)

func TestAtomic(t *testing.T) {
	var x int32 = 0
	group := sync.WaitGroup{}

	for i := 0; i < 1000; i++ {

		go func() {
			group.Add(1)
			for j := 0; j < 100; j++ {
				atomic.AddInt32(&x, 1)
			}
			group.Done()
		}()

	}

	group.Wait()
	fmt.Println("Counter : ", x)

}
