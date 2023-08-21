package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var cache = map[int]Book{}

var rnd = rand.New(rand.NewSource(time.Now().UnixNano()))

func main() {

	waitGroup := sync.WaitGroup{}
	mutx := sync.RWMutex{}

	cacheCh := make(chan Book)
	dbCh := make(chan Book)

	for i := 0; i < 10; i++ {
		id := rnd.Intn(10) + 1
		waitGroup.Add(2)
		go func(id int, wg *sync.WaitGroup, mx *sync.RWMutex, ch chan Book) {
			cacheBook, isInCache := queryCache(id, mx)

			if isInCache {
				fmt.Println("from cache")
				fmt.Println(cacheBook)
			}
			wg.Done()
		}(id, &waitGroup, &mutx, cacheCh)

		go func(id int, wg *sync.WaitGroup, mx *sync.RWMutex, ch chan Book) {
			dbBook, isInDb := queryDatabase(id, mx)

			if isInDb {
				fmt.Println("from db")
				fmt.Println(dbBook)
			}
			wg.Done()
		}(id, &waitGroup, &mutx, dbCh)
	}

	fmt.Print("hi")
	waitGroup.Wait()
}

func queryCache(id int, mx *sync.RWMutex) (Book, bool) {
	mx.RLock()
	book, ok := cache[id]
	mx.RUnlock()
	return book, ok
}

func queryDatabase(id int, mx *sync.RWMutex) (Book, bool) {

	for _, book := range books {
		if book.ID == id {
			mx.Lock()
			cache[book.ID] = book
			mx.Unlock()
			return book, true
		}
	}
	return Book{}, false
}
