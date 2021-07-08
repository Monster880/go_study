package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generator() chan int{
	out := make(chan int)
	go func() {
		for{
			time.Duration(time.Sleep(rand.Intn(1500)) * time.Millisecond))
		}
	}()
	return out
}

func main(){
	var c1, c2 chan int
	//for{
		select {
		case n := <- c1:
			fmt.Println("Received from c1:",n)
		case n:= <- c2:
			fmt.Println("Received from c2:",n)
		default:
			fmt.Println("No value received")
		}
	//}
}