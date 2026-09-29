package main

import (
	"fmt"
	"math/rand"
	"time"
)

func processNum(numChan chan int) {
	for num := range numChan {
		fmt.Println("processing number", num)
		time.Sleep(time.Second)
	}
}

func sum(result chan int, num1 int, num2 int) {
	numResult := num1 + num2
	result <- numResult
}

func task(done chan bool) {
	defer func() {
		done <- true
	}()

	fmt.Println("Processing...")
}

// make the channel receive only and send only
func emailSender(emailChan <-chan string, doneEmail chan<- bool) {
	defer func() {
		doneEmail <- true
	}()
	for email := range emailChan {
		fmt.Println("Sending email to", email)
		time.Sleep(time.Second)
	}
}

func main() {
	emailChan := make(chan string, 100)

	emailChan <- "ayush@example.com"
	emailChan <- "ankit@example.com"

	fmt.Println(<-emailChan)
	fmt.Println(<-emailChan)

	done := make(chan bool)
	go task(done)

	<-done // block

	// ===================================================================

	doneEmail := make(chan bool)

	go emailSender(emailChan, doneEmail)
	for i := range 5 {
		emailChan <- fmt.Sprintf("%d@gmail.com", i)
	}

	fmt.Println("done sending")

	// this is important
	close(emailChan)
	<-doneEmail

	// ====================================================================

	// Multiple channel

	chan1 := make(chan int)
	chan2 := make(chan string)

	go func() {
		chan1 <- 10
	}()

	go func() {
		chan2 <- "ping"
	}()

	// Match the loop count with the number of sends otherwise deadlock
	for range 2 {
		select {
		case chan1Val := <-chan1:
			fmt.Println("Received data from channel 1", chan1Val)
		case chan2Val := <-chan2:
			fmt.Println("Reveived data from channel 2", chan2Val)
		}
	}

	// ====================================================================

	result := make(chan int)

	go sum(result, 4, 5)

	res := <-result

	fmt.Println(res)

	numChan := make(chan int)

	go processNum(numChan)

	for {
		numChan <- rand.Intn(100)
	}

	// messageChannel := make(chan string)
	// messageChannel <- "ping"  // blocking
	// msg := <- messageChannel
	// fmt.Println(msg)
}
