package main

import (
	"fmt"
	"github.com/zencoder/disque-go/disque"
	"golang.org/x/net/context"
	"log"
	"math/rand"
	"sync"
	"time"
)

var p *disque.DisquePool

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

const queueName = "myqueue"

var msgLength int
var src []rand.Source

// RandStringBytesMaskImprSrc retruns fast random string
// See: http://stackoverflow.com/questions/22892120/how-to-generate-a-random-string-of-a-fixed-length-in-golang
func RandStringBytesMaskImprSrc(randSource, n int) string {
	b := make([]byte, n)
	s := src[randSource]
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, s.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = s.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}
	return string(b)
}

func init() {
	hosts := []string{"127.0.0.1:7711"} // array of 1 or more Disque servers
	cycle := 1000                       // check connection stats every 1000 Fetch's
	capacity := 10                      // initial capacity of the pool
	maxCapacity := 50                   // max capacity that the pool can be resized to
	idleTimeout := 15 * time.Minute     // timeout for idle connections
	p = disque.NewDisquePool(hosts, cycle, capacity, maxCapacity, idleTimeout)
}

// Put gets a randSoruce idx to ensure some thread safety
func Put(randSource int) {
	var d *disque.Disque
	var err error
	d, err = p.Get(context.Background()) // get a connection from the pool
	if err != nil {
		log.Println(err)
	}

	_, err = d.Push(queueName, RandStringBytesMaskImprSrc(randSource, msgLength), time.Second*20)
	if err != nil {
		log.Fatal(err)
	}
	p.Put(d) // return a connection to the pool
}

// Get gets from given queue
func Get() {
	var d *disque.Disque
	var err error
	d, err = p.Get(context.Background()) // get a connection from the pool
	if err != nil {
		log.Println(err)
	}

	_, err = d.Fetch(queueName, time.Second*1)
	if err != nil {
		log.Fatal(err)
	}
	p.Put(d) // return a connection to the pool

}

func PutTest(threads, N int) {
	var wg sync.WaitGroup
	start := time.Now()

	for j := 0; j < threads; j++ {
		wg.Add(1)
		a := rand.NewSource(time.Now().UnixNano())
		src = append(src, a)

		go func(idx int) {
			defer wg.Done()
			for i := 0; i < N; i++ {
				Put(idx)
			}
		}(j)
	}

	wg.Wait()
	duration := time.Since(start)
	total := N * threads
	speed := float64(total) / duration.Seconds()

	fmt.Printf("[SET] %12s %8d %7.3f/s\n", duration, total, speed)
}

func GetTest(threads, N int) {
	var wg sync.WaitGroup
	start := time.Now()

	for j := 0; j < threads; j++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			for i := 0; i < N; i++ {
				Get()
			}
		}()
	}

	wg.Wait()
	duration := time.Since(start)
	total := N * threads
	speed := float64(total) / duration.Seconds()

	fmt.Printf("[GET] %12s %8d %7.3f/s\n", duration, total, speed)
}

func main() {
	Ns := []int{10000, 20000, 40000}
	messageLens := []int{1, 4, 8, 16, 20, 200, 1000}
	ThreadCounts := []int{1, 2, 4, 8, 10, 16}

	for _, n := range Ns {
		for _, threads := range ThreadCounts {
			for _, messageLen := range messageLens {
				msgLength = messageLen
				
				fmt.Printf("Msg: %4d N: %7d Threads: %4d \n", msgLength, n, threads)
				PutTest(threads, n)
				GetTest(threads, n)
				
				fmt.Println("======================================")

				time.Sleep(time.Second * 1)

			}
		}
	}
}
