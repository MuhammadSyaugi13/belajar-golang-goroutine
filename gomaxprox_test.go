package belajar_golang_goroutine

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

func TestGomaxprox(t *testing.T) {

	group := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		group.Add(1)
		go func() {
			time.Sleep(3 * time.Second)
			group.Done()
		}()
	}

	totalCPU := runtime.NumCPU()
	fmt.Println("total cpu : ", totalCPU)

	runtime.GOMAXPROCS(20) // merubah jumlah thread
	totalThread := runtime.GOMAXPROCS(-2)
	fmt.Println("total thread : ", totalThread)

	totalGoroutine := runtime.NumGoroutine()
	fmt.Println("total goroutine : ", totalGoroutine)

	group.Wait()

}
