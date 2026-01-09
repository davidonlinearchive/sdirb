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

	// GOBRUTE HEADER
	fmt.Println("===============================================================")
	fmt.Println("Gobrute (BETA)")
	fmt.Println("By davidonlinearchive")
	fmt.Println("===============================================================")

	if err := cfg.runGoBrute(); err != nil {
		fmt.Println("Error: ", err)
	}

}
