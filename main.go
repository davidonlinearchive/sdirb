package main

import (
	"flag"
	"fmt"
	"os"
)

type config struct {
	targetURL    string
	wordlistPath string
	threads      int
	timeout      int
}

func main() {
	cfg := &config{}
	flag.StringVar(&cfg.targetURL, "u", "", "Target URL")
	flag.StringVar(&cfg.wordlistPath, "w", "", "Path to wordlist file")
	flag.IntVar(&cfg.threads, "t", 20, "Number of threads")
	flag.IntVar(&cfg.timeout, "timeout", 5, "Timeout in seconds")

	flag.Parse()

	if cfg.targetURL == "" || cfg.wordlistPath == "" {
		fmt.Println("Error: URL and wordlist required")
		flag.Usage()
		os.Exit(1)
	}

	//SDirB HEADER
	fmt.Println("===============================================================")
	fmt.Println("sdirb v1.0")
	fmt.Println("By @davidonlinearchive")
	fmt.Println("===============================================================")
	fmt.Printf("URL:            			%s\n", cfg.targetURL)
	fmt.Printf("HTTP Method:				GET\n")
	fmt.Printf("Threads:				%d\n", cfg.threads)
	fmt.Printf("Wordlist:				%s\n", cfg.wordlistPath)
	fmt.Printf("Timeout:				%ds\n", cfg.timeout)
	fmt.Println("===============================================================")

	if err := cfg.runDirBrute(); err != nil {
		fmt.Println("Error: ", err)
	}

}
