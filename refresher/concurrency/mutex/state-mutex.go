package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

func main()  {
	var state = make(map[int]int)
	var mutex = &sync.Mutex{}

	var readOperations uint64
	var writeOperations uint64
	total := 0

	for r := 0; r < 100; r++ {
		go func()  {
			for {
				key := rand.Intn(5)
				mutex.Lock()
				total += state[key]
				mutex.Unlock()
				atomic.AddUint64(&readOperations, 1)

				time.Sleep(time.Microsecond)
			}
		}()
	}

	for w:= 0; w < 10; w++ {
		go func()  {
			for {
				key := rand.Intn(5)
				value := rand.Intn(5)
				mutex.Lock()
				state[key] = value
				mutex.Unlock()
				atomic.AddUint64(&writeOperations, 1)

				time.Sleep(time.Millisecond)
			}
		}()
	}
	
	time.Sleep(time.Second)

	readOperationsFinal := atomic.LoadUint64(&readOperations)
	fmt.Println("read operations:" , readOperationsFinal)
	writeOperationsFinal := atomic.LoadUint64(&writeOperations)
	fmt.Println("read operations:" , writeOperationsFinal)
	fmt.Println("total:" , total)

	mutex.Lock()
	fmt.Println("state", state)
	mutex.Unlock()

}