// package main

// import (
// 	"fmt"
// 	"io"
// 	"net/http"
// 	"sync"
// 	"time"
// )

// func fetchURL(url string, ch chan<- string, wg *sync.WaitGroup) {
// 	defer wg.Done() // Signal when the goroutine is done

// 	start := time.Now()
// 	resp, err := http.Get(url)
// 	if err != nil {
// 		ch <- fmt.Sprintf("Error fetching %s: %s", url, err)
// 		return
// 	}
// 	defer resp.Body.Close()

// 	_, _ = io.Copy(io.Discard, resp.Body) // Read the entire body
// 	secs := time.Since(start).Seconds()
// 	ch <- fmt.Sprintf("%.2fs elapsed fetching %s", secs, url)
// }

// func main() {
// 	urls := []string{
// 		"https://www.google.com",
// 		"https://www.facebook.com",
// 		"https://www.amazon.com",
// 	}

// 	ch := make(chan string)
// 	var wg sync.WaitGroup

// 	for _, url := range urls {
// 		wg.Add(1) // Increment the wait group counter
// 		go fetchURL(url, ch, &wg)
// 	}

// 	go func() {
// 		wg.Wait() // Wait for all goroutines to finish
// 		close(ch) // Close the channel to signal completion
// 	}()

// 	for msg := range ch {
// 		fmt.Println(msg)
// 	}
// }