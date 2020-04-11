package main

import (
	"bufio"
	"flag"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
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

	// received file from server
	resp, err := http.Get(inputUrl)
	if err != nil {
		log.Fatal("Download error: ", err)
	}
	defer resp.Body.Close()

	// create output file
	f, err := os.Create(outputFile)
	if err != nil {
		log.Fatal("Download error: ", err)
	}
	defer f.Close()

	// read server response line by line
	scanner := bufio.NewScanner(resp.Body)
	i := 0
	for scanner.Scan() {
		l := scanner.Text()

		// if line contains url address
		if strings.HasPrefix(l, "http") {
			// download file part
			part, err := downloadFilePart(l)
			if err != nil {
				log.Fatal("Download part error: ", err)
			}

			// write part to output file
			if _, err = f.Write(part); err != nil {
				log.Fatal("Write part to output file: ", err)
			}
			log.Printf("Download part %d\n", i)
			i++
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

//downloadFilePart download file part from server
func downloadFilePart(url string) ([]byte, error) {
	result := make([]byte, 0)

	resp, err := http.Get(url)
	if err != nil {
		return result, err
	}

	if result, err = ioutil.ReadAll(resp.Body); err != nil {
		return result, err
	}

	return result, err
}
