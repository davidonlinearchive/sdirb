# sdirb (Simple DirB).

A fast, concurrent directory brute-forcing tool written in Go. `sdirb` helps security professionals discover hidden directories and files on web servers during penetration testing and security assessments.

## Features

- Concurrent scanning with configurable thread pools
- Real-time progress tracking with visual progress bar
- Clean, organized output showing status codes and URLs

## Installation

### Prerequisites

- Go 1.25 or higher


### Install
```bash
git clone https://github.com/davidonlinearchive/sdirb.git
cd sdirb
go build
```

## Usage

Basic usage example:
```bash
./sdirb -u http://example.com -w wordlist.txt
```

### Options
```
-u string
    Target URL (required)
-w string
    Path to wordlist file (required)
-t int
    Number of threads (default: 20)
-timeout int
    Timeout in seconds (default: 5)
```

### Examples
```bash
# Basic scan
./sdirb -u http://example.com -w wordlist.txt

# Scan with 50 threads
./sdirb -u http://example.com -w wordlist.txt -t 50

# Custom timeout
./sdirb -u http://example.com -w wordlist.txt -timeout 10
```

## Output
```
[200] http://example.com/admin
[403] http://example.com/private
Progress  15% |██░░░░░░░░░░░░| (1500/10000) [45s:4m30s]
```

## Disclaimer
This tool is for educational and ethical testing purposes only. Always obtain proper authorization before testing any systems you do not own.
