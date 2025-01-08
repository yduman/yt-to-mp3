# yt-to-mp3

Read a list of YouTube URLs from a text file, then use goroutines to concurrently download and convert each video to MP3. Store all resulting MP3 files in a designated output folder.

## Requirements

- [`ffmpeg`](https://www.ffmpeg.org/)
- [`yt-dlp`](https://github.com/yt-dlp/yt-dlp)

The input file is expected to have one link for each line. Did not test playlist links.

```txt
<link 1>
<link 2>
<link 3>
...
```

## Usage

Build the binary with Go

```console
cd cmd && go build -o yt2mp3
```

Run the binary

```console
./yt2mp3 <filepath to urls> <output path> [concurrency]
```

`concurrency` is by default `10`. It controls how many downloads to run in parallel. The higher the number, the more speed up but also more network/CPU usage.