all: build

build:
	CGO_ENABLED=1 go build -ldflags -H=windowsgui -o m3u8-downloader

win:
	CGO_ENABLED=1 GOOS=windows GOARCH=amd64 CC=x86_64-w64-mingw32-gcc go build -ldflags -H=windowsgui -o m3u8-downloader.exe