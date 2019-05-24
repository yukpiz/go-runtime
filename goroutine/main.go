package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	n := 100
	fmt.Printf("CPUs: %d\n", runtime.NumCPU())
	fmt.Printf("GOMAXPROCS: %d\n", runtime.GOMAXPROCS(1))
	fmt.Printf("GOMAXPROCS: %d\n", runtime.GOMAXPROCS(10))
	fmt.Printf("GOMAXPROCS: %d\n", runtime.GOMAXPROCS(1))

	var wg sync.WaitGroup
	wg.Add(3)

	go func() {
		defer wg.Done()
		for i := 0; i < n; i++ {
			fmt.Printf("a%d\n", i)
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < n; i++ {
			fmt.Printf("b%d\n", i)
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < n; i++ {
			fmt.Printf("c%d\n", i)
		}
	}()

	wg.Wait()
}
