package main

import (
	"fmt"

	"github.com/kimeunse/golang/multi"
	"github.com/kimeunse/golang/single"
)

func main() {
	multi.Mul()
	strr, err := single.Hello("")
	if err == nil {
		fmt.Println(strr)
		fmt.Println("main")
	} else {
		fmt.Println(err)
	}
}
