package main

import (
	"fmt"
	"sync"
	"time"
)

// worker function
func workerNode(id int, jobChan <-chan int, resChan chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	for value := range jobChan {
		msg := fmt.Sprintf("Worker %d got %d", id, id+value)
		time.Sleep(time.Duration(id) * time.Second)
		resChan <- msg

	}

}

func main() {
	var noJobs = 10
	var noWorkers = 3
	var wg sync.WaitGroup
	jobChan := make(chan int, noJobs)
	resChan := make(chan string, noJobs)

	for i := range noJobs {
		jobChan <- (i + 1)
	}
	close(jobChan)
	for i := range noWorkers {
		wg.Add(1)
		go workerNode((i + 1), jobChan, resChan, &wg)
	}
	go func() {
		wg.Wait()
		close(resChan)
	}()
	for msg := range resChan {
		fmt.Println(msg)
	}

}
