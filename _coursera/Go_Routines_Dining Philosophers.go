package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Phil struct {
	n           int
	left, right *sync.Mutex
}

var finished = make(chan int)
var perm = make(chan int, 2)
var wg sync.WaitGroup

func host() {
	defer wg.Done()

	perm <- 0
	perm <- 0
	for i := 0; i < N*3-2; i++ {
		perm <- <-finished
	}
	<-finished
	<-finished
}

func (p *Phil) eat() {
	defer wg.Done()

	for i := 0; i < 3; i++ {
		<-perm
		if rand.Intn(2) == 0 {
			p.left.Lock()
			p.right.Lock()
		} else {
			p.right.Lock()
			p.left.Lock()
		}

		fmt.Printf("starting to eat %d\n", p.n)
		time.Sleep(100 * time.Millisecond)
		fmt.Printf("finishing eating %d\n", p.n)

		p.left.Unlock()
		p.right.Unlock()

		finished <- 0
	}
}

const (
	N = 5
)

func main() {
	var cs = make([]sync.Mutex, N, N)
	var phils = make([]Phil, N, N)

	for i := 0; i < N; i++ {
		phils[i].n = i
		phils[i].left = &cs[i]
		phils[i].right = &cs[(i+1)%5]
	}

	wg.Add(N + 1)

	go host()

	for i := 0; i < N; i++ {
		go phils[i].eat()
	}

	wg.Wait()
}
