package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan string)

	go msg1(ch1)
	go msg2(ch2)

	for {
		if ch1 == nil && ch2 == nil {
			break
		}
		select {
		case val, ok := <-ch1:
			if !ok {
				ch1 = nil
				continue
			}
			fmt.Println("ch1:", val)
		case val, ok := <-ch2:
			if !ok {
				ch2 = nil
				continue
			}
			fmt.Println("ch2:", val)
		default:
			time.Sleep(50 * time.Millisecond)
		}
	}
}

func msg1(ch chan int) {
	defer func() {
		p := recover()
		if p != nil {
			fmt.Println("panic", p)
		}
		close(ch)
	}()
	for i := 0; i < 10; i++ {
		num := rand.Intn(5)
		if num == 0 {
			panic("Zero")
		}
		ch <- num
	}
}

func msg2(ch chan string) {
	for i := 0; i < 10; i++ {
		ch <- "hello"
	}
	close(ch)
}
