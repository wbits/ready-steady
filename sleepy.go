package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)
	for i := 1; i < 6; i++ {
		go sleepyGopher(i, c)
	}
	timeout := time.After(2 * time.Second)
	for i := 0; i < 5; i++ {
		select {
		case gopherID := <-c:
			fmt.Println("gopher ", gopherID, " has finished sleeping")
		case <-timeout:
			fmt.Println("my patience ran out")
			return
		}
	}
}

func sleepyGopher(id int, c chan int) {
	fmt.Println("... ", id, " snore ...")
	c <- id
}
