```go
package main

import (

        "fmt"
        "sync"
)

func main() {
        var wg sync.WaitGroup
        var count int
        var mu sync.Mutex // Add a mutex
        for i := 0; i < 10; i++ {
                wg.Add(1)
                go func() {
                        defer wg.Done()
                        mu.Lock() // Lock the mutex
                        count++
                        mu.Unlock() // Unlock the mutex
                }()
        }
        wg.Wait()
        fmt.Println(count)
}
```