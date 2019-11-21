package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	fmt.Println("Welcome")
	defer fmt.Println("Bye") // will called in the end

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

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

	go func() {
		s := <-signals
		fmt.Println("Catch termination signal:", s)
		forever <- false
	}()

	<-forever
}
