package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	count := 10
	squaredChan := make(chan int)

	// Запускаем первую горутину
	go func() {
		rawChan := make(chan int)

		// Запускаем вторую горутину внутри первой
		go func() {
			for n := range rawChan {
				squaredChan <- n * n
			}
			close(squaredChan)
		}()

		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		for i := 0; i < count; i++ {
			n := r.Intn(101)
			rawChan <- n
		}
		close(rawChan)
	}()

	// В main читаем из squaredChan
	fmt.Println("Squares:")
	for sq := range squaredChan {
		fmt.Printf("%d ", sq)
	}

}
