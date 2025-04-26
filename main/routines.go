package main

import (
	"fmt"
	"sync"
	"time"
)

var m = sync.Mutex{}
var rwm = sync.RWMutex{} // has separate read and write locks as well
var wg = sync.WaitGroup{}
var dbData = []string{"id1", "id2", "id3", "id4", "id5"}
var results = []string{}

func routines() {
	fmt.Println("Go Routines!")

	// concurrency != parallel execution
	// concurrency --> having multiple tasks in progress at the same time (could be back and forth)
	// parallel exec --> have simultaneous exec due to multiple cpu cores

	t0 := time.Now()
	for i := 0; i < len(dbData); i++ {
		wg.Add(1)
		go dbCall(i) // won't wait for dbCall to finish
	}

	wg.Wait()
	fmt.Printf("Total execution time: %v \n", time.Since(t0))
	fmt.Printf("The results are: %v \n", results)
	// using goroutines directly can lead to corupt data as 2 processes might write to the same mem location
	// at the same time! use mutex to make concurrency safe

}

// simulates a db call
func dbCall(i int) {
	var delay float32 = 2000
	time.Sleep(time.Duration(delay) * time.Millisecond)
	fmt.Printf("The result from the db is %v \n", dbData[i])
	// m.Lock() // the position of the lock can totally fuck up your concurrency! be mindful
	// results = append(results, dbData[i])
	// m.Unlock() // unless the mutex is unlocked by the current thread, other threads cannot exec the above line
	save(dbData[i])
	log()
	wg.Done()
}

func save(result string) {
	rwm.Lock() // all locks including rlocks must be released first
	results = append(results, result)
	rwm.Unlock()
}

func log() {
	rwm.RLock() // multiple routines can hold read locks
	fmt.Printf("The current results are %v \n", results)
	rwm.RUnlock()
}

// func dbCall(i int) {
// 	var delay float32 = rand.Float32() * 2000
// 	time.Sleep(time.Duration(delay) * time.Millisecond)
// 	fmt.Printf("The result from the db is %v \n", dbData[i])
// 	wg.Done()
// }

/*

NOTE:- If a go routine is computationally expensive (ex- incrementing i=0-->10000000) then the execution time
depends on the number of cores available in the cpu, but if it is not so computationally expensive (ex- a db call that
takes time) then it is a constant time operation no matter the number of concurrent calls made!

*/
