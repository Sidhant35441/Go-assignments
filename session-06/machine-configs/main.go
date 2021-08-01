package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("logical processor :", runtime.NumCPU())
	fmt.Println("operating system :", runtime.GOOS)
	fmt.Println("system architecture :", runtime.GOARCH)
	fmt.Println("max processes :", runtime.GOMAXPROCS(1))

}
