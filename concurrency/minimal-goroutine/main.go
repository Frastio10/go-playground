package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)

	go func() {
		fmt.Println("Running slow ass function...")
		time.Sleep(2 * time.Second)

		fmt.Println("Done")
		ch <- "mwahahahah"
	}()
	fmt.Println("Waiting for the function to run...")

	data := <-ch

	fmt.Println("recv", data)
}
