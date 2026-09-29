package main

import (
	"fmt"
	"sync"
)

func task(id int, w *sync.WaitGroup) {
	defer w.Done()
	fmt.Println(id, &w)
}

func main() {
	var waitGroup sync.WaitGroup

	for i := 0; i <= 10; i++ {
		waitGroup.Add(1)
		go task(i, &waitGroup)
	}

	waitGroup.Wait()
}
