package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	jobs := make(chan int, 10)

	// HTTP Requests example
	urls := []string{
		"https://www.google.com",
		"https://www.yahoo.com",
		"http://www.localhost.com",
	}
	urlJobs := make(chan string, len(urls))

	// Start three workers goroutines
	for w := 1; w <= 3; w++ {
		wg.Add(1)
		go worker(w, jobs, &wg)
		// go readUrl(w, urlJobs, &wg)
	}

	for w2 := 4; w2 <= 5; w2++ {
		wg.Add(1)
		go readUrl(w2, urlJobs, &wg)
	}

	// send 10 jobs to channel
	for j := 1; j <= 10; j++ {
		jobs <- j
		fmt.Println("Job send:", j)
	}
	close(jobs)

	for _, url := range urls {
		urlJobs <- url
	}
	close(urlJobs)

	// wait for all workers to be finish
	wg.Wait()
	fmt.Println("All jobs are completed")

}

func worker(w int, jobs <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		fmt.Printf("Worker %d is processing job %d\n", w, job)
		time.Sleep(time.Second)
	}
}

func readUrl(w int, urlJobs <-chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for url := range urlJobs {
		response, err := http.Get(url)
		if err != nil {
			fmt.Printf("Woker %d fails to fetch url %s: %v\n", w, url, err)
			continue
		}
		fmt.Printf("Worker %d fetch the url %s with the status: %s\n", w, url, response.Status)
		response.Body.Close()
	}
}
