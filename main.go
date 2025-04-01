package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

func main() {
	start := time.Now()

	const n = 10

	var wg sync.WaitGroup
	wg.Add(n)

	for i := 0; i <= n; i++ {
		go func() {
			defer wg.Done()
			resp, err := http.Get("https://google.com")

			if err != nil {
				panic(err)
			}

			defer resp.Body.Close()
			fmt.Println("ok")
		}()
	}

	wg.Wait()
	fmt.Println(time.Since(start))

}
