package main

import (
	"fmt"
	"math/rand/v2"
)

func initSlice(out chan<- int) {
	defer close(out)
	for range 10 {
		out <- rand.IntN(101)
	}
}

func square(in <-chan int, out chan<- int) {
	defer close(out)
	for v := range in {
		out <- v * v
	}
}

func main() {
	nums := make(chan int)
	squares := make(chan int)

	go initSlice(nums)
	go square(nums, squares)

	for v := range squares {
		fmt.Println(v)
	}
}
