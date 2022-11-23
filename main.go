package main

import (
	"fmt"
	"time"
)

func f(n int) {
	for i := 0; i < 10; i++ {
		fmt.Println(n, ":", i)
		time.Sleep(100)
	}
}
func main() {
	for i := 0; i < 10; i++ {
		go f(i)
	}

	fmt.Println("Waiting...")
	var input string
	result, err := fmt.Scanln(&input)
	if err != nil {
		fmt.Println("Something wrong...")

	}
	fmt.Println(result)
}
