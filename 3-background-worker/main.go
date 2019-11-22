package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type BackgroundWorker interface {
	Run()
	Terminate()
}

// ----------

func NewHelloWriterWorker() BackgroundWorker {
	return &HelloWriterWorker{
		terminationChan:     make(chan bool, 0),
		terminationDoneChan: make(chan bool, 0),
	}
}

type HelloWriterWorker struct {
	terminationChan     chan bool
	terminationDoneChan chan bool
}

func (hw *HelloWriterWorker) Run() {
	ticker := time.NewTicker(2 * time.Second)

	// make decision - do work or terminate
	for {
		select {
		case <-ticker.C:
			fmt.Println("Hello")
		case <-hw.terminationChan:
			hw.terminationDoneChan <- true
			return
		}
	}
}

func (hw *HelloWriterWorker) Terminate() {
	fmt.Println("HelloWriterWorker stopping...")
	hw.terminationChan <- true
	fmt.Println("HelloWriterWorker is stopped.")
	<-hw.terminationDoneChan
}

// ----------

func NewWorldWriterWorker() BackgroundWorker {
	return &WorldWriterWorker{
		terminationChan:     make(chan bool, 0),
		terminationDoneChan: make(chan bool, 0),
	}
}

type WorldWriterWorker struct {
	terminationChan     chan bool
	terminationDoneChan chan bool
}

func (ww *WorldWriterWorker) Run() {
	ticker := time.NewTicker(2 * time.Second)

	// make decision - do work or terminate
	for {
		select {
		case <-ticker.C:
			fmt.Println("World")
		case <-ww.terminationChan:
			ww.terminationDoneChan <- true
			return
		}
	}
}

func (ww *WorldWriterWorker) Terminate() {
	fmt.Println("WorldWriterWorker stopping...")
	ww.terminationChan <- true
	fmt.Println("WorldWriterWorker is stopped.")
	<-ww.terminationDoneChan
}

// ----------

func main() {
	fmt.Println("Welcome to 3-background-worker sample...")
	defer fmt.Println("Bye") // will called in the end

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	forever := make(chan bool)

	workerOne := NewHelloWriterWorker()
	go workerOne.Run()
	defer workerOne.Terminate()

	workerTwo := NewWorldWriterWorker()
	go workerTwo.Run()
	defer workerTwo.Terminate()

	go func() {
		s := <-signals
		fmt.Println("Catch termination signal:", s)
		forever <- false
	}()

	<-forever
}
