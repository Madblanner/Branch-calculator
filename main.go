package main

import (
	"fmt"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)
	ch4 := make(chan int)
	go Add(2, 2, ch1)
	go Add(3, 6, ch2)
	go Multip(7, 7, ch3)
	go Del(9, 3, ch4)
	i := <-ch1 + <-ch2 + <-ch3 + <-ch4
	fmt.Println("Numbers of examples =", i)

}

func Add(x, y int, ch chan int) {
	i := x + y
	fmt.Printf("%d + %d = %d\n", x, y, i)
	ch <- 1

}
func Multip(x, y int, ch chan int) {
	i := x * y
	fmt.Printf("%d * %d = %d\n", x, y, i)
	ch <- 1

}
func Del(x, y int, ch chan int) {
	i := x / y
	fmt.Printf("%d / %d = %d\n", x, y, i)
	ch <- 1
}
