package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"sync"

	"github.com/yduman/yt-to-mp3/internal/downloader"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Printf("Usage: %s <urlsFile> <outputFolder> [concurrency]\n", filepath.Base(os.Args[0]))
		os.Exit(1)
	}

	filePath := os.Args[1]
	out := os.Args[2]
	concurrency := 10

	if len(os.Args) >= 4 {
		if val, err := strconv.ParseInt(os.Args[3], 10, 64); err == nil && val > 0 {
			concurrency = int(val)
		}
	}

	urls, err := downloader.ReadLinks(filePath)
	if err != nil {
		log.Fatalf("Failed to read file: %v\n", err)
	}

	if err := os.MkdirAll(out, 0755); err != nil {
		log.Fatalf("Failed to create out folder: %v\n", err)
	}

	sem := make(chan struct{}, concurrency)
	var wg sync.WaitGroup

	for _, url := range urls {
		wg.Add(1)
		sem <- struct{}{}

		go func(u string) {
			defer wg.Done()
			defer func() { <-sem }()

			if err := downloader.ToMP3(u, out); err != nil {
				log.Printf("Error downloading %s: %v\n", u, err)
			}
		}(url)
	}

	wg.Wait()
	fmt.Println("All links done.")
}
