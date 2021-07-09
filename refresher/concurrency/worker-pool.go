package main

import (
	"fmt"
	"time"
)

func cWorker(id int, tasks <-chan int, results chan <- int) {
	for t := range tasks {
		fmt.Println("worker", id, "started task", t)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished task", t)
		results <- t * 2
	}
}

func main() {
	const numTasks = 5
	tasks := make(chan int, numTasks)
	results := make(chan int, numTasks)

	for a := 1; a <= 3; a++ {
		go cWorker(a, tasks, results)
	}

	for b := 1; b <= numTasks; b++ {
		tasks <- b
	}
	close(tasks)

	for c := 1; c <= numTasks; c++ {
		<- results
	}
}