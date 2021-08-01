package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func printEvenNumbers() {
	defer wg.Done()
	fmt.Println("Even Numbers :")
	for i := 1; i <= 20; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
}
func printOddNumbers() {
	defer wg.Done()
	fmt.Println("Odd Numbers :")
	for i := 1; i <= 20; i++ {
		if i%2 != 0 {
			fmt.Println(i)
		}
	}
}

func main() {
	fmt.Println("get in main function..")
	fmt.Println("Hello GoLang")
	wg.Add(3)   //adding 2 go rountine to execute
	go func() { //annonymous function
		fmt.Println("Annonymous func..")
		wg.Done()
	}()
	go printEvenNumbers() //using go routine
	go printOddNumbers()
	fmt.Println("get out main function..")
	wg.Wait() //wait till all the go rountines completed
}
