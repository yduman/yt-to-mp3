package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"sync"
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

	urls, err := readLinks(filePath)
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
			defer func() { <-sem }() // release slot when done

			if err := toMP3(u, out); err != nil {
				log.Printf("Error downloading %s: %v\n", u, err)
			}
		}(url)
	}

	wg.Wait()
	fmt.Println("All links done.")
}

func toMP3(url, outDir string) error {
	// yt-dlp --extract-audio --audio-format mp3 -o <outputDir>/%(title)s.%(ext)s <URL>
	cmd := exec.Command("yt-dlp", "--extract-audio", "--audio-format", "mp3", "-o", filepath.Join(outDir, "%(title)s.%(ext)s"), url)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func readLinks(path string) ([]string, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer f.Close()

	var urls []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		urls = append(urls, line)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return urls, nil
}
