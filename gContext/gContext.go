package gcontext

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func Go_context_withCancel() {
	wg.Add(1)
	ctx, cancel := context.WithCancel(context.Background())

	go PrintTick(ctx)

	time.Sleep(5 * time.Second)
	cancel()
	wg.Wait()

}

func Go_context_withDeadline() {
	wg.Add(1)

	d := time.Now().Add(3 * time.Second)
	ctx, cancel := context.WithDeadline(context.Background(), d)

	go PrintTick(ctx)

	time.Sleep(5 * time.Second)
	cancel()
	wg.Wait()

}

func PrintTick(ctx context.Context) {
	tick := time.Tick(time.Second)
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Done", ctx.Err())
			wg.Done()
			return
		case <-tick:
			fmt.Println("tick")
		}
	}
}
