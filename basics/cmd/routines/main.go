package main

import (
	"fmt"
	"sync"
	"time"
)

var m = sync.Mutex{}
var rm = sync.RWMutex{}
var wg = sync.WaitGroup{}
var dbData = []string{"id1", "id2", "id3", "id4", "id5"}
var results = []string{}

func main() {
	start := time.Now()
	// for i := 0; i < len(dbData); i++ {
	// 	wg.Add(1)
	// 	go dbCall(i)
	// }
	// for i := 0; i < len(dbData); i++ {
	// 	wg.Add(1)
	// 	go dbCallR(i)
	// }
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		// go dbCallS(i)
		go count()
	}
	wg.Wait()
	fmt.Printf("Total execution time: %v\n", time.Since(start))
	// fmt.Printf("The results are: %v\n", results)
}

func dbCall(i int) {
	// simulate DB call delay
	var delay float32 = 2000
	time.Sleep(time.Duration(delay) * time.Millisecond)
	fmt.Println("The result from the database is:", dbData[i])
	m.Lock()
	results = append(results, dbData[i])
	m.Unlock()
	wg.Done()
}

func dbCallR(i int) {
	// simulate DB call delay
	var delay float32 = 2000
	time.Sleep(time.Duration(delay) * time.Millisecond)
	save(dbData[i])
	log()
	wg.Done()
}

func save(result string) {
	rm.Lock()
	results = append(results, result)
	rm.Unlock()
}

func log() {
	rm.RLock()
	fmt.Printf("The current results are: %v\n", results)
	rm.RUnlock()
}

func dbCallS(i int) {
	// simulate DB call delay
	var delay float32 = 2000
	time.Sleep(time.Duration(delay) * time.Millisecond)
	wg.Done()
}

func count() {
	var res int
	for i := 0; i < 100_000_000; i++ {
		res += 1
	}
	wg.Done()
}
