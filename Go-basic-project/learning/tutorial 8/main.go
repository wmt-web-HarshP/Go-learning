package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var n = sync.RWMutex{}
var wg = sync.WaitGroup{}
var dbData = []string{"id1", "id2", "id3", "id4", "id5"}
var results = []string{}

func main() { 
	t0 := time.Now()
	for i := 0; i < len(dbData); i++ {
		wg.Add(1)
		go dbCall(i)
	}
	wg.Wait()
	fmt.Printf("\nTotal execution time:%v", time.Since(t0))
	fmt.Printf("\nThe result are%v", results)
}

func dbCall(i int) {
	var delay float32 = rand.Float32() * 2000
	time.Sleep(time.Duration(delay) * time.Millisecond)
	fmt.Println("The result from the database is:", dbData[i])
	save(dbData[i])
	log()
	// n.Lock()
	// result = append(result, dbData[i])
	// n.Unlock()
	wg.Done()
}

func save(result string) {
	n.Lock()
	results = append(results, result)
	n.Unlock()
}

func log() {
	n.RLock()
	fmt.Printf("\nThe current results are%v", results)
	n.RUnlock()
}