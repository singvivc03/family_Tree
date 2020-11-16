package main

import (
	"fmt"
)

func main() {
	c := make(chan int) // unbuffered channel of integers
	go func() {
		sendMessage("Hi")
		c <- 1
	}()
	sendMessage("Vivek")
	<-c
}

func sendMessage(message string) {
	fmt.Println(message)
}
