package cli

import (
	"fmt"
)

const bufferSize = 10000

type Downloader struct {
	logger     chan string
	downloadFn func(string, string, chan string) error
}

func (c *Downloader) Download(url string, file string) error {
	c.logger <- fmt.Sprintf("Start download %s", url)
	err := c.downloadFn(url, file, c.logger)
	if err != nil {
		return err
	}
	c.logger <- fmt.Sprintf("File saved %s", file)

	return nil
}

func New(downloader func(string, string, chan string) error) *Downloader {
	d := Downloader{
		logger:     make(chan string, bufferSize),
		downloadFn: downloader,
	}
	go func() {
		for {
			m := <-d.logger
			fmt.Println(m)
		}
	}()
	return &d
}
