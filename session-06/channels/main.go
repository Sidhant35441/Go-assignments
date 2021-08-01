package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	wg.Add(10)
	ch := make(chan int, 10)
	go printNum1(ch)
	fmt.Println(<-ch) //receive from channel
	go printNum2(ch)
	fmt.Println(<-ch)
	go printNum3(ch)
	fmt.Println(<-ch)
	go printNum4(ch)
	fmt.Println(<-ch)
	go printNum5(ch)
	fmt.Println(<-ch)
	go printNum6(ch)
	fmt.Println(<-ch)
	go printNum7(ch)
	fmt.Println(<-ch)
	go printNum8(ch)
	fmt.Println(<-ch)
	go printNum9(ch)
	fmt.Println(<-ch)
	go printNum10(ch)
	fmt.Println(<-ch)
	fmt.Println("main func")
	wg.Wait()
}
func printNum1(ch chan int) {
	defer wg.Done()
	ch <- 10 //send to channel
}
func printNum2(ch chan int) {
	defer wg.Done()
	ch <- 20
}
func printNum3(ch chan int) {
	defer wg.Done()
	ch <- 30
}
func printNum4(ch chan int) {
	defer wg.Done()
	ch <- 40
}
func printNum5(ch chan int) {
	defer wg.Done()
	ch <- 50
}
func printNum6(ch chan int) {
	defer wg.Done()
	ch <- 60
}
func printNum7(ch chan int) {
	defer wg.Done()
	ch <- 70
}
func printNum8(ch chan int) {
	defer wg.Done()
	ch <- 80
}
func printNum9(ch chan int) {
	defer wg.Done()
	ch <- 90
}
func printNum10(ch chan int) {
	defer wg.Done()
	ch <- 100
}
