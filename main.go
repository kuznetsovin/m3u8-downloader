package main

import (
	"fmt"
	"github.com/kuznetsovin/m3u8-downloader/cli"
	"github.com/kuznetsovin/m3u8-downloader/downloader"
	"github.com/kuznetsovin/m3u8-downloader/gui"
	"os"
)

func printHelpCommand() {
	fmt.Println("See information: m3u8-downloader help")
	os.Exit(0)
}

func printCliHelpCommand() {
	fmt.Println("Usage: m3u8-downloader cli help")
	os.Exit(0)
}

func main() {
	argsCount := len(os.Args)
	if argsCount < 2 {
		printHelpCommand()
	}

	mode := os.Args[1]
	switch mode {
	case "gui":
		a := gui.New(downloader.Download)
		a.Run()
	case "cli":
		switch argsCount {
		case 2:
			printCliHelpCommand()
		case 3:
			if os.Args[2] == "help" {
				fmt.Println("Usage: m3u8-downloader cli [url] [output file]")
				os.Exit(0)
			}
			printCliHelpCommand()
		case 4:
			url := os.Args[2]
			outputFile := os.Args[3]
			c := cli.New(downloader.Download)
			if err := c.Download(url, outputFile); err != nil {
				fmt.Printf("Download error: %v\n", err)
				os.Exit(1)
			}
		default:
			printHelpCommand()
		}
	case "help":
		fmt.Println("m3u8-downloader support some modes:")
		fmt.Println("	gui - run gui app")
		fmt.Println("	cli - run command-line interface. More information: m3u8-downloader cli help")
	default:
		printHelpCommand()
	}

}
