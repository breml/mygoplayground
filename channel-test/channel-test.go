package main

import "time"

func do1(c chan int) {
	defer close(c)
	for i := 0; i < 5; i++ {
		c <- i
		time.Sleep(1 * time.Second)
	}
}

func main() {
	println("Hello")
	c1 := make(chan int, 0)
	go do1(c1)
	for j := 0; j < 20; j++ {
		v, ok := <-c1
		if !ok {
			break
		}
		println("Run:", v)
	}
}
