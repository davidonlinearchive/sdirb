package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"
)

func (c *config) runDirBute() error {
	jobs := make(chan string, c.threads)
	var wg sync.WaitGroup

	client := &http.Client{
		Timeout: time.Duration(c.timeout) * time.Second,
	}

	// open wordlist file
	file, err := os.Open(c.wordlistPath)
	if err != nil {
		return fmt.Errorf("Error opening wordlist: %w", err)
	}
	defer file.Close()

	// start worker pool
	for i := 0; i < c.threads; i++ {
		//  wg.Go simplfies traditional goroutines management (added in go 1.25)
		wg.Go(func() {
			for path := range jobs {
				u := fmt.Sprintf("%s/%s", strings.TrimRight(c.targetURL, "/"),
					strings.TrimLeft(path, "/"))
				resp, err := client.Get(u)
				if err == nil {
					if resp.StatusCode != 404 {
						fmt.Printf("[%d] %s\n", resp.StatusCode, u)
					}
					resp.Body.Close()
				}
			}
		})
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		jobs <- scanner.Text() // converts Paths to strings and sends it to a jobs channel
	}
	if err := scanner.Err(); err != nil {
		return err
	}
	close(jobs)

	wg.Wait() // Ensures waitgroup counter is zero before exiting
	return nil
}
