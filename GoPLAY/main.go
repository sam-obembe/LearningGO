package main

import (
	"demo/dicts"
	"fmt"
	"math"
	"math/rand"
	"sync"
)

func main() {
	//sliceWork()
	//makingMaps()
	//structPlay()
	//panicRecover()
	//learnprimitives.SimpleDataTypes()
	//webserver.RunWebServer()

	//myRoutinesAndWaits()

	//concurrentLooping()

	//messwithstrings.TryStringFormatting("Sam")
	dicts.DictDabble()

}

func arithmetic() {
	randomNumber := rand.Int()
	fmt.Printf("random number = %d", randomNumber)

	root := math.Sqrt(4)
	fmt.Printf("Square root of 4 is %f \n", root)

	a := 2
	b := 3
	addition := Add(a, b)
	fmt.Printf("The sum of %d and %d is %d \n", a, b, addition)
}

// can also be written as Add(x,y int)int{...}
func Add(x int, y int) int {
	return x + y
}

func multiReturn(name, surname string) (string, string) {
	return surname, name
}

func dabbleWithConstants() {
	const address = 13505

	const name string = "Samuel"

	const c = address

	const (
		howdi = "hi"
		adios = "bye" //if a value isn't provided, it copies the recent value in the group
	)

	const h = iota //would have value of 0
}

func playWithPointers() {
	a := "hello pointer"
	b := &a

	fmt.Println(b)  //memory address
	fmt.Println(*b) //hello pointer

	*b = "subarashi pointer"

	fmt.Println(b)  //memory address
	fmt.Println(*b) //subarashi pointer

	b = new(string)
	fmt.Println(b)  //new memory address
	fmt.Println(*b) //empty string

	*b = "issall over"
	fmt.Println(*b)
}

// i.e dictionaries
func makingMaps() {
	var myMap map[string]int // map[keyType]valueType

	myMap = map[string]int{"foo": 1, "bar": 3}
	fmt.Println(myMap)

	myMap["foo"] = 99
	fmt.Println(myMap)

	delete(myMap, "bar")
	fmt.Println(myMap)

	val, ok := myMap["bar"] //accessing a key that doesn't exist
	fmt.Println(val, ok)

	var complexMap map[string][]string
	complexMap = map[string][]string{
		"sam":  {"coffee", "golang"},
		"dave": {"shayo", "java"},
	}

	fmt.Println(complexMap)
}

func structPlay() {
	type myStruct struct {
		name string
		id   int
	}

	var me myStruct
	me = myStruct{name: "Sam", id: 27}

	fmt.Println(me)
}

func infiniteLoop() {
	i := 1
	for {
		fmt.Println(i)
		i++

		if i > 95 {
			break
		}
	}
}

func conditionalLoop() {
	i := 1

	for i <= 3 {
		fmt.Println(i)
		i++
	}

}

func regularForLoop() {
	for i := 1; i < 3; i++ {
		fmt.Println(i)
	}
}

func loopingThroughCollections() {
	mySlice := []string{"Hey", "you", "there"}

	for ki, val := range mySlice {
		fmt.Println(ki, val)
	}

	myMap := map[string]int{"age": 27, "height": 170}

	for ki, val := range myMap {
		fmt.Println(ki, val)
	}
}

// if -else
func ifElse() {
	myNum := 25

	if myNum > 20 {
		fmt.Println("gt 25")
	} else if myNum > 15 {
		fmt.Println("gt 15")
	} else {
		fmt.Println("lt 25")
	}

	//alternative, using initializer
	if myAltNum := 25; myAltNum > 20 {
		fmt.Println("my alt num:", "gt 25")
	} else if myAltNum > 15 {
		fmt.Println("my alt num:", "gt 15")
	} else {
		fmt.Println("my alt num:", "lt 25")
	}
}

// switch statements
func switchCase() {
	mySwitchVal := 10

	switch mySwitchVal {
	case 1:
		fmt.Println("tis 1")
	case 5 * 2, 4 * 5:
		fmt.Println("tis multiple of 10")
	default:
		fmt.Println("default")
	}

	//logical switch can also be written as switch i:=50;{} . Go implies the test case is true by default
	switch logical := 8; true {
	case logical < 5:
		fmt.Println("less than 5")
	case logical > 5:
		fmt.Println("greater than 5")
	default:
		fmt.Println("unaccounted for")
	}

}

// deferred functions
func deferredFuncs() {
	fmt.Print("main 1")

	fmt.Println("defer 1")

	fmt.Println("main 2")

	fmt.Println("defer 2")
	// output main 1, main2, defer 2, defer 1
}

// panic and recovery
func panicRecover() {
	fmt.Println("Starting")
	woulPanic()
	fmt.Println("ending")

	dividend, divisor := 10, 0
	res := divideToPanic(dividend, divisor)
	fmt.Printf("Dividing %v by %v = %v\n", dividend, divisor, res)

	dividend, divisor = 10, 5
	res = divideToPanic(dividend, divisor)
	fmt.Printf("Dividing %v by %v = %v\n", dividend, divisor, res)

}

func woulPanic() {

	defer func() {
		fmt.Println(recover())
	}()
	fmt.Println("going to panic soon")

	panic("yepaaaa")

	//fmt.Println("continuing after panic")
}

func divideToPanic(dividend, divisor int) int {
	defer func() {
		msg := recover()
		if msg != nil {
			fmt.Println(msg)
		}
	}()
	return dividend / divisor
}

// goto statements
func gotoStatements() {

}

// go rorutines and wait groups
func myRoutinesAndWaits() {
	var wg sync.WaitGroup
	ch := make(chan string)
	wg.Add(1)

	go func() {
		ch <- "this message"
	}()

	go func() {
		msg := <-ch
		fmt.Println(msg)
		wg.Done()
	}()

	wg.Wait()
}

//select statements

func selectStatements() {
	ch1, ch2 := make(chan string), make(chan string)

	go func() {
		ch1 <- "message to channel 1"
	}()

	go func() {
		ch2 <- "message to channel 2"
	}()

	select {
	case msg := <-ch1:
		fmt.Println(msg)
	case msg := <-ch2:
		fmt.Println(msg)
	default:
		fmt.Println("no messages available")
	}
}

func concurrentLooping() {
	ch := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		close(ch)
	}()

	for msg := range ch {
		fmt.Println(msg)
	}
}
