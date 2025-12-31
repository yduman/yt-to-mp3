# yt-to-mp3

Go CLI tools for downloading YouTube videos and converting them to MP3 using concurrent downloads.

## Requirements

- [`ffmpeg`](https://www.ffmpeg.org/)
- [`yt-dlp`](https://github.com/yt-dlp/yt-dlp)

## Build

```console
go build -o yt2mp3 ./cmd/yt2mp3
go build -o yt2mp3-playlist ./cmd/yt2mp3-playlist
```

## Usage

### Batch Mode (`yt2mp3`)

Download multiple videos from a text file containing URLs (one per line).

```console
./yt2mp3 <urls-file> <output-dir> [concurrency]
```

**Example:**
```console
./yt2mp3 links.txt ./music 5
```

**Input file format:**
```txt
https://www.youtube.com/watch?v=VIDEO_ID_1
https://www.youtube.com/watch?v=VIDEO_ID_2
https://www.youtube.com/watch?v=VIDEO_ID_3
```

### Playlist Mode (`yt2mp3-playlist`)

Download all videos from a YouTube playlist.

```console
./yt2mp3-playlist <playlist-url> <output-dir> [concurrency]
```

**Example:**
```console
./yt2mp3-playlist "https://www.youtube.com/playlist?list=PLxxxxx" ./music 5
```

## Options

| Argument | Description | Default |
|----------|-------------|---------|
| `concurrency` | Number of parallel downloads | 10 |

Higher concurrency increases speed but also network/CPU usage.
