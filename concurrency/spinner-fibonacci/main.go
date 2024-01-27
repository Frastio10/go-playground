package main

import (
	"fmt"
	"time"
)

func main() {
	go spinner(1 * time.Second) // this basically will run forever until the execution time is finished
	const n = 45
	fibN := fib(n) // slow asf

	fmt.Printf("\rFibonacci(%d) = %d\n", n, fibN)
}

func spinner(delay time.Duration) {
	for {
		fmt.Println("is running....")
		for _, r := range `-\|/` { // to print a spinner...
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}

func fib(x int) int {
	if x < 2 {
		return x
	}

	return fib(x-1) + fib(x-2)
}
