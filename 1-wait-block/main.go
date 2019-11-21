package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome")
	defer fmt.Println("Bye")

	forever := make(chan bool)

	go func() {
		for {
			fmt.Println("Hello")
			time.Sleep(2 * time.Second)
		}
	}()

	go func() {
		for {
			fmt.Println("World")
			time.Sleep(2 * time.Second)
		}
	}()

	<-forever
}
