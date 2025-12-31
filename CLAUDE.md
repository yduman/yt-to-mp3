# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

Go CLI tools for downloading YouTube videos and converting them to MP3 using `yt-dlp` and `ffmpeg`.

## Build & Run

```bash
# Build both binaries
go build -o yt2mp3 ./cmd/yt2mp3
go build -o yt2mp3-playlist ./cmd/yt2mp3-playlist

# Batch mode (from URL file)
./yt2mp3 <urls-file> <output-dir> [concurrency]

# Playlist mode
./yt2mp3-playlist <playlist-url> <output-dir> [concurrency]
```

- `urls-file`: Text file with one YouTube URL per line
- `playlist-url`: YouTube playlist URL
- `output-dir`: Directory for MP3 output (created if doesn't exist)
- `concurrency`: Optional, defaults to 10 parallel downloads

## Architecture

```
cmd/
  yt2mp3/main.go          # Batch download from URL file
  yt2mp3-playlist/main.go # Download from playlist URL
internal/
  downloader/
    downloader.go         # Shared: ToMP3(), ReadLinks()
    playlist.go           # ExtractPlaylistURLs()
```

Both binaries use a semaphore pattern (buffered channel + WaitGroup) for concurrent downloads.

## External Dependencies

Requires `yt-dlp` and `ffmpeg` installed on the system (not Go dependencies).
