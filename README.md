# m3u8-downloader

Downloader m3u8 files from URL.

## Install

1. Go get:
```bash
go get -v github.com/kuznetsovin/m3u8-downloader
```

2. Install from source:

```bash
git clone https://github.com/kuznetsovin/m3u8-downloader.git
cd m3u8-downloader
go build -o m3u8-downloader
```

## How to use

```bash
m3u8-downloader -url <path_to_m3u8> -file <path_to_output_file>
```

## Update for v2

Version 2 use [fyne](https://fyne.io) for GUI and now doesn't support console mode.