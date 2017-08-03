package main

import (
	"fmt"
	"time"
	"math/rand"
)

func Goroutine(channel chan bool) {
	rand_number := rand.Intn(5)

	fmt.Printf("goroutine waiting for %d seconds\n", rand_number)

	time.Sleep(time.Second * time.Duration(rand_number))

	channel <- true
}

func main() {
	complete_channel := make(chan bool)

	go Goroutine(complete_channel)

	<-complete_channel
}