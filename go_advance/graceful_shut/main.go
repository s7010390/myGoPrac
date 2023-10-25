package main

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(3)
	myTime := time.Now()
	go doSometing("go", wg)
	go doSometing("rust", wg)
	go doSometing("reat", wg)

	wg.Wait()
	fmt.Println("Endtime:", time.Since(myTime))

}

func worker(s string, wg *sync.WaitGroup) {
	defer wg.Done()

	b := sha256.Sum256([]byte(s))
	fmt.Println(b)
	fmt.Println(
		base64.StdEncoding.EncodeToString(b[:]),
	)
}

func doSometing(s string, wg *sync.WaitGroup) {
	defer wg.Done()

	time.Sleep(time.Millisecond * 100)
	fmt.Println(s)
}
