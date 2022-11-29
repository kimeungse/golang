package gchannel

import (
	"fmt"
	"os"
	"time"
)

func counter() {
	for i := 0; ; i++ {
		time.Sleep(time.Second)
		ch <- i
	}
}

func stop() {
	time.Sleep(time.Second * 10)
	quit <- true
}

var ch chan int = make(chan int)
var quit chan bool = make(chan bool)

func Go_channel() {
	go counter()
	go stop()

	for {
		select {
		case count := <-ch:
			fmt.Println(count)
		case <-quit:
			os.Exit(1)
		}
	}
}
