## sync.WaitGroup

The difference between `wg.Add(numWorkers)` and `wg.Add(1)` within a loop is in the number of times you increment the `sync.WaitGroup` counter, and it affects how many goroutines you expect to wait for before allowing the main goroutine to proceed.

1. `wg.Add(numWorkers)`:
    - When you use `wg.Add(numWorkers)` outside the loop, you are incrementing the `WaitGroup` counter by `numWorkers` all at once. This implies that you expect `numWorkers` goroutines to call `wg.Done()` to signal their completion before the `WaitGroup` can proceed.

    - If you are starting multiple goroutines concurrently (e.g., in a loop), and you want to wait for all of them to complete, you should use `wg.Add(numWorkers)` where `numWorkers` is the total number of goroutines you are starting. This is useful when you know the exact number of goroutines in advance.

    - Example:
      ```go
      var wg sync.WaitGroup
 
      numWorkers := 3
      wg.Add(numWorkers)
 
      for i := 0; i < numWorkers; i++ {
          go worker(i, &wg)
      }
      
      wg.Wait()
      ```

2. `wg.Add(1)`:
    - When you use `wg.Add(1)` inside the loop, you are incrementing the `WaitGroup` counter by 1 for each iteration of the loop. This implies that you expect each iteration of the loop to start a single goroutine, and you want to wait for each of these goroutines to complete individually before the `WaitGroup` can proceed.

    - This approach is useful when you are starting goroutines dynamically within a loop, and you want to wait for each goroutine to finish before moving on to the next one.

    - Example:
      ```go
      var wg sync.WaitGroup
 
      numWorkers := 3
 
      for i := 0; i < numWorkers; i++ {
          wg.Add(1)
          go worker(i, &wg)
      }
      
      wg.Wait()
      ```

In summary, the choice between `wg.Add(numWorkers)` and `wg.Add(1)` depends on your specific use case and whether you are starting a fixed number of goroutines in advance or dynamically within a loop.
