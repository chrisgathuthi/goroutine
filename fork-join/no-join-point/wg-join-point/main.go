package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		work()

	}()
	wg.Wait()
	fmt.Println("done waiting, main exits")
	// join
}
func work() {
	time.Sleep(500 * time.Millisecond)
	fmt.Println("Printing some stuff")

}
