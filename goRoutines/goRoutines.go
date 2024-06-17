package main

import (
	"fmt"
	"sync"
)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done() //Signal that this goRoutine is done
	fmt.Printf("Worker %d starting\n", id)
	//simulate some work
	fmt.Printf("worker %d done\n", id)
}

func main() {
	fmt.Println("Will learn about Concurrency")

	var wg sync.WaitGroup

	//Start 3 worker goRoutines
	for i := 1; i < 3; i++ {
		wg.Add(1) //Increment the WaitGroup Counter
		go worker(i, &wg)
	}

	//wait all workers to Finish
	wg.Wait()
	fmt.Println("All workers done")
}
