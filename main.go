package main

import (
	"github.com/kuznetsovin/m3u8-downloader/downloader"
	"github.com/kuznetsovin/m3u8-downloader/gui"
)

func main() {
	a := gui.New(downloader.Download)
	a.Run()
}
