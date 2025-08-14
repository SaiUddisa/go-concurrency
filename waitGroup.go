package main

import (
	"fmt"
	"sync"
	"time"
)

func sayHai(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Start  this is %d", id)
	time.Sleep(2 * time.Second)
	fmt.Println("End this is %d", id)
}
func main() {
	// runtime.GOMAXPROCS(2)
	var wg sync.WaitGroup
	//adding one go routine
	wg.Add(10)
	go sayHai(1, &wg)
	go sayHai(2, &wg)
	go sayHai(3, &wg)
	go sayHai(4, &wg)
	go sayHai(5, &wg)
	go sayHai(6, &wg)
	go sayHai(7, &wg)
	go sayHai(8, &wg)
	go sayHai(9, &wg)
	go sayHai(10, &wg)
	wg.Wait()
	fmt.Println("this is the main function")

}
