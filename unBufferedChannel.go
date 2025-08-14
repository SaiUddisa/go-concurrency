package main

import (
	"fmt"
	"sync"
)

func worker(id int, ch chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	msg := fmt.Sprintf("Worker %d Reporting", id)
	ch <- msg
}
func main() {
	var wg sync.WaitGroup
	var ch = make(chan string)
	wg.Add(5)
	go worker(1, ch, &wg)
	go worker(2, ch, &wg)
	go worker(3, ch, &wg)
	go worker(4, ch, &wg)
	go worker(5, ch, &wg)

	go func() {
		wg.Wait()
		close(ch)
	}()
	for msg := range ch {
		fmt.Println(msg)
	}

	fmt.Println("\nWorker Attendence Done")
}
