package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 2)
	exit := make(chan int)

	go func() {
		for i := 0; i < 5; i++ {

			fmt.Println(time.Now(), i, "Sending")
			ch <- i
			fmt.Println(time.Now(), i, "Sent")
			time.Sleep(1 * time.Second)
		}

		fmt.Println(time.Now(), "All tasks are completed. closing channel..")
		close(ch)
	}()

	go func() {
		for {
			select {
			case v, open := <-ch: // double assignment, the second assignment is  an indicator whether the channel is active or not
				if !open {
					// there are many ways to do it, put a value and then discard it, or just close the channel entirely. (maybe?)
					// otherwise, there will be a deadlock
					close(exit)
					// exit <- 1
					return
				}

				fmt.Println(time.Now(), "Recv", v)
			}
		}

		// Another cool way to do it, basically it will loop based on the value that is received, and print whatever data is received, only works on a single channel though
		// for v := range ch {
		// fmt.Println(time.Now(), "received", v)
		// }

	}()

	fmt.Println(time.Now(), "Waiting...")

	<-exit // it blocks until the channel received a value or it's closed.

	fmt.Println(time.Now(), "Exiting")

}
