package main

import (
	"fmt"
	"time"
)

func main() {
	var start time.Time = time.Now()
	for i := 0; i <= 100000000000; i++ {
		
	}
	var end time.Time = time.Now();

	fmt.Println(end.Sub(start))
}