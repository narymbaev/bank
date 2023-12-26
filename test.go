package main

import (
	"fmt"
	"runtime"
	"time"
)

func synchronousFunction() {
	fmt.Println("Synchronous function started")
	//time.Sleep(2 * time.Second)
	fmt.Println("Synchronous function completed")
}
//wg sync.WaitGroup
func prtFunction() error {
	//defer wg.Done()
	//fmt.Println("Function started")
	//time.Sleep(2 * time.Second)
	//fmt.Println("Function completed")
	return nil
}

func main() {
	// Synchronous function
	synchronousFunction()

	runtime.GOMAXPROCS(runtime.NumCPU())

	// Asynchronous function using Goroutine
	//var wg sync.WaitGroup
	//wg.Add(1)
	errs := make(chan error)
	for i := 0; i < 5; i++ {
		go func(a int) {
			err := prtFunction()
			fmt.Printf("process %v started \n", a)
			if a % 2 == 1 {
				time.Sleep(7 * time.Second)
			} else {
				time.Sleep(3 * time.Second)
			}

			errs <- err
			fmt.Printf("process %v ended \n", a)
		}(i)
		fmt.Println("loop end")
	}

	time.Sleep(2*time.Second)

	//fmt.Println("After sleep")

	for i := 0; i < 5; i++ {
		fmt.Println("BLOCKING CHANNEL")
		err := <- errs
		time.Sleep(2 * time.Second)
		fmt.Println("RELEASING CHANNEL")
		if err != nil {
			fmt.Println("error detected")
		}
	}
	// Main function continues execution without waiting for Goroutines to finish
	fmt.Println("Main function continues while Goroutine is running")

	// Wait for the Goroutine to finish
	//wg.Wait()

	fmt.Println("Main function completed")
}
