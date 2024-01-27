// this code will cause a deadlock cause there are no goroutines running atm
package main

import "fmt"

func main() {
	var ch chan int
	ch = make(chan int)

	ch <- 10

	v := ch

	fmt.Println("recv", v)
}
