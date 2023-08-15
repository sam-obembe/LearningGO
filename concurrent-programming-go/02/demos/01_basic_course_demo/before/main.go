package main

import (
	"fmt"
	"math/rand"
	"time"
)

var cache = map[int]Book{}

var rnd = rand.New(rand.NewSource(time.Now().UnixNano()))

func main() {

	for i := 0; i < 10; i++ {
		id := rnd.Intn(10) + 1

		cacheBook, isInCache := queryCache(id)

		if isInCache {
			fmt.Println("from cache")
			fmt.Println(cacheBook)
			continue
		}

		dbBook, isInDb := queryDatabase(id)

		if isInDb {
			fmt.Println("from db")
			fmt.Println(dbBook)
			continue
		}

		fmt.Printf("Book with id  not found %v", id)
		time.Sleep(150 * time.Millisecond)
	}
	fmt.Print("hi")
}

func queryCache(id int) (Book, bool) {
	book, ok := cache[id]
	return book, ok
}

func queryDatabase(id int) (Book, bool) {
	time.Sleep(100 * time.Millisecond)
	for _, book := range books {
		if book.ID == id {
			cache[book.ID] = book
			return book, true
		}
	}
	return Book{}, false
}
