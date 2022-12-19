package gcontext

//https://dev-yakuza.posstree.com/ko/golang/context/
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

func Go_WithTimeout() {
	wg.Add(1)
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	go PrintTick(ctx)
	time.Sleep(time.Second * 5)
	cancel()
	wg.Wait()
}

func Go_WithValue() {
	wg.Add(1)
	ctx := context.WithValue(context.Background(), "v", 3)
	go square(ctx)
	wg.Wait()
}

func square(ctx context.Context) {
	if v := ctx.Value("v"); v != nil {
		n := v.(int)
		fmt.Println("Square:", n*n)
	}
	wg.Done()
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
