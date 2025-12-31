package downloader

import (
	"bufio"
	"os"
	"os/exec"
	"path/filepath"
)

// ToMP3 downloads a YouTube video and converts it to MP3.
func ToMP3(url, outDir string) error {
	cmd := exec.Command("yt-dlp", "--extract-audio", "--audio-format", "mp3", "-o", filepath.Join(outDir, "%(title)s.%(ext)s"), url)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

// ReadLinks reads URLs from a file, one per line, skipping empty lines.
func ReadLinks(path string) ([]string, error) {
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
