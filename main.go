package main

import (
	"fmt"
)

func main() {
	ch := make(chan string)
	go Add(2, 2, ch)
	go Add(3, 6, ch)
	go Multip(7, 7, ch)
	go Del(9, 3, ch)
	fmt.Println(<-ch)
}

func Add(x, y int, ch chan string) {
	i := x + y
	fmt.Printf("%d + %d = %d\n", x, y, i)
	ch <- "exit"
}
func Multip(x, y int, ch chan string) {
	i := x * y
	fmt.Printf("%d * %d = %d\n", x, y, i)
	ch <- "exit"
}
func Del(x, y int, ch chan string) {
	i := x / y
	fmt.Printf("%d / %d = %d\n", x, y, i)
	ch <- "exit"
}
