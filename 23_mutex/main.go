package main

import (
	"fmt"
	"sync"
)

type post struct {
	views int
	mu    sync.Mutex
}

func (p *post) inc(weightGroup *sync.WaitGroup) {
	defer func() {
		p.mu.Unlock()
		weightGroup.Done()
	}()

	p.mu.Lock()
	p.views += 1
}

func main() {
	var weightGroup sync.WaitGroup

	myPost := post{views: 0}

	for range 10000 {
		weightGroup.Add(1)
		go myPost.inc(&weightGroup)
	}

	weightGroup.Wait()
	fmt.Println(myPost.views)
}
