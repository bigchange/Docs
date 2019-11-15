package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

var endpoints = []string{
	"100.69.62.1:3232",
	"100.69.62.32:3232",
	"100.69.62.42:3232",
	"100.69.62.81:3232",
	"100.69.62.11:3232",
	"100.69.62.113:3232",
	"100.69.62.101:3232",
}

// var indexes = []int{0, 1, 2, 3, 4, 5, 6}

// 重点在这个 shuffle
func shuffle(indexes []int) {
	// b := rand.Perm(n)
	for i := len(indexes); i > 0; i-- {
		lastIdx := i - 1
		idx := rand.Intn(i)
		indexes[lastIdx], indexes[idx] = indexes[idx], indexes[lastIdx]
	}
}
func request() error {
	var err error
	var indexes = []int{0, 1, 2, 3, 4, 5, 6}
	shuffle(indexes)
	maxRetryTimes := 3
	fmt.Println(indexes)
	idx := 0
	for i := 0; i < maxRetryTimes; i++ {
		// err = apiRequest(params, indexes[idx])
		if err == nil {
			break
		}
		idx++
	}

	if err != nil {
		// logging
		return err
	}

	return nil
}
func main() {
	var n sync.Map
	// sync.Mutex
	// sync.Pool
	// sync.WaitGroup
	n.Store("key", "value")
	fmt.Println(EditDistance("ALGORITHM", "ALTRUISTIC"))
	request()
}
