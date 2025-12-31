package downloader

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

// ExtractPlaylistURLs extracts all video URLs from a YouTube playlist.
func ExtractPlaylistURLs(playlistURL string) ([]string, error) {
	cmd := exec.Command("yt-dlp", "--flat-playlist", "--print", "url", playlistURL)

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		return nil, fmt.Errorf("yt-dlp failed: %w\nstderr: %s", err, stderr.String())
	}

	var urls []string
	for _, line := range strings.Split(stdout.String(), "\n") {
		line = strings.TrimSpace(line)
		if line != "" {
			urls = append(urls, line)
		}
	}

	if len(urls) == 0 {
		return nil, fmt.Errorf("no URLs extracted from playlist")
	}

	return urls, nil
}
