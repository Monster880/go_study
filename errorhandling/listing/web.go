package main

import (
	"fmt"
	"time"
)

func worker(id int, c chan int){
	for{
		fmt.Printf("Worker %d received %d\n", id, <-c)
	}
}

func CreateWorker(id int)chan<- int{
	c := make(chan int)
	go worker(id, c)
	return  c
}

func chanDemo(){
	var channels [10] chan<- int
	for i := 0; i<10; i++{
		channels[i] = CreateWorker(i)
	}

	for i :=0; i<10; i++{
		channels[i] <- 'a' + i
	}

	for i :=0; i<10; i++{
		channels[i] <- 'A' + i
	}
	time.Sleep(time.Millisecond)
}

func bufferChannel(){
	c := make(chan int, 3)
	go worker(0, c)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	c <- 'd'
	time.Sleep(time.Millisecond)
}

func channelClose(){
	c := make(chan int)
	go worker(0, c)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	c <- 'd'
	close(c)
	time.Sleep(time.Millisecond)
}

func main() {
	// chanDemo()
	//bufferChannel()
	channelClose()
}
