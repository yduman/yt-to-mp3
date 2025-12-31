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
		fmt.Printf("Usage: %s <playlist-url> <outputFolder> [concurrency]\n", filepath.Base(os.Args[0]))
		os.Exit(1)
	}

	playlistURL := os.Args[1]
	out := os.Args[2]
	concurrency := 10

	if len(os.Args) >= 4 {
		if val, err := strconv.ParseInt(os.Args[3], 10, 64); err == nil && val > 0 {
			concurrency = int(val)
		}
	}

	fmt.Printf("Extracting URLs from playlist: %s\n", playlistURL)
	urls, err := downloader.ExtractPlaylistURLs(playlistURL)
	if err != nil {
		log.Fatalf("Failed to extract playlist URLs: %v\n", err)
	}
	fmt.Printf("Found %d videos in playlist\n", len(urls))

	if err := os.MkdirAll(out, 0755); err != nil {
		log.Fatalf("Failed to create out folder: %v\n", err)
	}

	fmt.Printf("Starting downloads with concurrency: %d\n", concurrency)

	sem := make(chan struct{}, concurrency)
	var wg sync.WaitGroup
	var failed int
	var mu sync.Mutex

	for _, url := range urls {
		wg.Add(1)
		sem <- struct{}{}

		go func(u string) {
			defer wg.Done()
			defer func() { <-sem }()

			if err := downloader.ToMP3(u, out); err != nil {
				log.Printf("Failed: %s - %v\n", u, err)
				mu.Lock()
				failed++
				mu.Unlock()
			}
		}(url)
	}

	wg.Wait()
	fmt.Printf("Completed: %d/%d successful\n", len(urls)-failed, len(urls))

	if failed > 0 {
		os.Exit(1)
	}
}
