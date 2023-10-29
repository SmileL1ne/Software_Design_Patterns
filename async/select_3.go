package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	cancelCh := make(chan bool)
	dataCh := make(chan int)

	go func(cancelCh chan bool, dataCh chan int) {
		val := 0
		for {
			select {
			case <-cancelCh:
				close(dataCh)
				return
			case dataCh <- val:
				val++
			}
		}
	}(cancelCh, dataCh)

	for curVal := range dataCh {
		fmt.Println("read", curVal)
		if curVal > 3 {
			fmt.Println("send cancel")
			cancelCh <- true
			// break
		}
	}
	time.Sleep(time.Second * 5)
	go func(cancelChan chan bool) {
		fmt.Println(<-cancelCh)
	}(cancelCh)
	cancelCh <- true
	time.Sleep(time.Second * 30)
	fmt.Println(runtime.NumGoroutine())

}
