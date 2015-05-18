package main
import "fmt"
func main() {
    start := time.Now()

    in := make(chan int)    // Channel on which work orders are received.
    out := make(chan []int) // Channel on which results are returned.
    go producer(in)
    go worker(in, out) // Launch one worker.
    consumer(out, 100)

    fmt.Println(time.Since(start))
}
func worker(in chan int, out chan []int) {
    for {
        order := <-in           // Receive a work order.
        result := factor(order) // Do some work.
        out <- result           // Send the result back.
    }
}
func producer(out chan int) {
    for order := 0; ; order++ {
        out <- order // Produce a work order.
    }
}

func consumer(in chan []int, n int) {
    for i := 0; i < n; i++ {
        result := <-in // Consume a result.
        fmt.Println("Consumed", result)
    }
}
