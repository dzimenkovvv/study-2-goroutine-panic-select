package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	generalChannel := make(chan int)
	numSources := 3

	done := make(chan struct{}, numSources)

	for i := 1; i <= numSources; i++ {
		go func(id int) {
			defer func() {
				p := recover()
				if p != nil {
					fmt.Println("")
					fmt.Println("Panic in source", id, "\nName:", p)
					fmt.Println("")
				}
				done <- struct{}{}
			}()

			for i := 0; i < 5; i++ {
				val := rand.Intn(5)
				if val == 0 {
					panic("Zero value")
				}
				generalChannel <- val
				time.Sleep(200 * time.Millisecond)
			}
		}(i)
	}
	var sum int

	active := numSources

	for active > 0 {
		select {
		case val := <-generalChannel:
			fmt.Println("Value:", val)
			sum += val

		case <-done:
			active--
		}
	}

	close(generalChannel)
	fmt.Println("----------------")
	fmt.Println("Total sum:", sum)
}
