package main

import (
	"fmt"
	"runtime"
	"sync"
	"unsafe"
)

type counter struct {
	val int
}

func (c *counter) Add(int) {
	c.val++
}

func (c *counter) Value() int {
	return c.val
}

func main() {
	runtime.GOMAXPROCS(2)

	var (
		wg    sync.WaitGroup
		mu    sync.Mutex
		meter counter
	)
	fmt.Printf("size of counter struct : %d\n", unsafe.Sizeof(meter))

	for i := 0; i < 1000; i++ {
		wg.Add(1)

		go func() {
			for j := 0; j < 1000; j++ {
				mu.Lock()
				meter.Add(1)
				mu.Unlock()
			}

			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println(meter.Value())
	fmt.Printf("size of counter struct : %d\n", unsafe.Sizeof(meter))
}
