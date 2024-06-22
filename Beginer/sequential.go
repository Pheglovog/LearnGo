package main

import (
	"fmt"
	"sync"
	_ "time"
)

func sequential() {
	var wg sync.WaitGroup
	done := make([]chan bool, 3)

	for i := range done {
		done[i] = make(chan bool)
	}

	data := []string{"one", "two", "three"}

	for i, v := range data {

		wg.Add(1)

		go func(n int, v string) {
			defer func() {
				wg.Done()
			}()

			x := <-done[n]

			fmt.Println(v)

			if n < 2 {
				done[n+1] <- x
			}
		}(i, v)
	}

	done[0] <- true

	wg.Wait()

	for i := range done {
		done[i] = make(chan bool)
	}
}
