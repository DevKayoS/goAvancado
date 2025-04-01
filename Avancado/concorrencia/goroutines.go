package concorrencia

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"
)

func goroutines() {
	start := time.Now()

	const n = 10

	var wg sync.WaitGroup
	wg.Add(n)

	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	for i := 0; i <= n; i++ {
		go func(ctx context.Context) {
			defer wg.Done()

			resp, err := http.Get("https://google.com")
			if err != nil {
				panic(err)
			}

			defer resp.Body.Close()
			fmt.Println("ok")
		}(ctx)
	}

	wg.Wait()
	fmt.Println(time.Since(start))

}
