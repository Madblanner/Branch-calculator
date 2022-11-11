package main

import (
	"fmt"
	"time"
)

func main() {
	go Add(2, 2)
	go Add(3, 6)
	go Multip(7, 7)
	go Del(9, 3)
	time.Sleep(1 * time.Second)
}

func Add(x, y int) {
	i := x + y
	fmt.Printf("%d + %d = %d\n", x, y, i)
}
func Multip(x, y int) {
	i := x * y
	fmt.Printf("%d * %d = %d\n", x, y, i)
}
func Del(x, y int) {
	i := x / y
	fmt.Printf("%d / %d = %d\n", x, y, i)
}
