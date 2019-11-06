package main
import (
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		HTTPServeMain() // http handler
		wg.Done()
	}()

	wg.Wait()
}
