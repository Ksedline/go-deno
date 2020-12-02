// https://levelup.gitconnected.com/use-go-channels-as-promises-and-async-await-ee62d93078ec
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func longRunningTask() <-chan int32 {
	r := make(chan int32)

	go func() {
		defer close(r)
		
		time.Sleep(time.Second * 3)
		r <- rand.Int31n(100)
	}()

	return r
}

func main() {
	r := <-longRunningTask()
	fmt.Println(r)
}
