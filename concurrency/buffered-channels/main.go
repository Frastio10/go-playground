package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	go func() {
		for i := 0; i < 3; i++ {
			fmt.Println(time.Now(), i, "Sending")
			ch <- i
			fmt.Println(time.Now(), i, "Sent")
		}

		fmt.Println(time.Now(), "All completed")
	}()

	time.Sleep(2 * time.Second)

	fmt.Println(time.Now(), "waiting for messages")

	fmt.Println(time.Now(), "recv: ", <-ch)
	fmt.Println(time.Now(), "recv: ", <-ch)
	fmt.Println(time.Now(), "recv: ", <-ch)

	fmt.Println("Exiting..")
}
