package main

import (
	"flag"
	"github.com/kuznetsovin/m3u8-downloader/downloader"
	"log"
)

func main() {

	var (
		inputUrl   string
		outputFile string
	)
	flag.StringVar(&inputUrl, "url", "", "url path to m3u8")
	flag.StringVar(&outputFile, "file", "", "path to output file")

	flag.Parse()
	if inputUrl == "" {
		log.Fatal("url is required")
	}

	if outputFile == "" {
		log.Fatal("file is required")
	}

	if err := downloader.Download(inputUrl, outputFile); err != nil {
		log.Fatal(err)
	}

}
