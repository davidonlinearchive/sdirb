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
	var mu sync.Mutex

	client := &http.Client{
		Timeout: time.Duration(c.timeout) * time.Second,
	}

	// open wordlist file
	file, err := os.Open(c.wordlistPath)
	if err != nil {
		return fmt.Errorf("Error opening wordlist: %w", err)
	}
	defer file.Close()

	// Count real paths (skipping comments) for the bar total
	scanner := bufio.NewScanner(file)
	totalPaths := 0
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line != "" && !strings.HasPrefix(line, "#") {
			totalPaths++
		}
	}
	if err := scanner.Err(); err != nil {
		return err
	}
	file.Seek(0, 0) // Reset to start of file

	bar := newProgressBar(totalPaths)

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
						mu.Lock()
						bar.Clear()
						fmt.Printf("[%d] %s\n", resp.StatusCode, u)
						mu.Unlock()
					}
					resp.Body.Close()
				}
				bar.Add(1)
			}
		})
	}

	scanner = bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		jobs <- line
	}

	if err := scanner.Err(); err != nil {
		return err
	}
	close(jobs)

	wg.Wait() // Ensures waitgroup counter is zero before exiting
	fmt.Println()
	return nil
}
