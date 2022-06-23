# m3u8-downloader

Downloader m3u8 files from URL. Downloader has 2 mode: `gui` and `cli`. For more information run `m3u8-downloader help`

## Install

1. Go get:
```bash
go install -v github.com/kuznetsovin/m3u8-downloader@latest
```

2. Install from source:

```bash
git clone https://github.com/kuznetsovin/m3u8-downloader.git
cd m3u8-downloader
make
```

## How to use

Run gui version:

```bash
./m3u8-downloader gui 
```

Run cli version:

```bash
./m3u8-downloader cli [url] [output_file] 
```

For more information run `m3u8-downloader help`
