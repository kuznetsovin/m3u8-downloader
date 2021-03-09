package downloader

import (
	"bufio"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

func Download(url, file string) error {
	// received file from server
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// create output file
	f, err := os.Create(file)
	if err != nil {
		return err
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
			part := make([]byte, 0)

			resp, err := http.Get(url)
			if err != nil {
				return err
			}

			if part, err = ioutil.ReadAll(resp.Body); err != nil {
				return err
			}

			if _, err = f.Write(part); err != nil {
				log.Fatal("Write part to output file: ", err)
			}
			log.Printf("Download part %d\n", i)
			i++
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return err
}
