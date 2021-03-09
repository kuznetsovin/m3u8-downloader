package gui

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func New(download func(string, string, chan string) error) fyne.App {
	a := app.New()
	w := a.NewWindow("m3u8 downloader")
	w.Resize(fyne.Size{Width: 500, Height: 200})

	urlInput := widget.NewEntry()
	outputInput := widget.NewEntry()

	logInfo := widget.NewTextGrid()
	logInfo.Resize(fyne.Size{Height: 100})

	downloaderState := make(chan string, 1000)
	go func() {
		for {
			msg := <-downloaderState
			logInfo.SetRow(len(logInfo.Rows)+1, textToGridRow(msg))
		}
	}()

	btn := widget.NewButton("Download", nil)
	btn.Resize(fyne.Size{Width: 80})
	btn.OnTapped = func() {
		btn.Disable()
		defer btn.Enable()

		url := urlInput.Text
		file := outputInput.Text
		downloaderState <- fmt.Sprintf("Start download %s", url)
		err := download(url, file, downloaderState)
		if err != nil {
			downloaderState <- err.Error()
		}
		downloaderState <- fmt.Sprintf("File saved %s", file)
	}

	w.SetContent(
		container.New(layout.NewGridLayout(1),
			container.New(layout.NewFormLayout(),
				widget.NewLabel("URL"),
				urlInput,
				widget.NewLabel("File"),
				outputInput,
			),
			container.NewScroll(logInfo),
			container.New(layout.NewCenterLayout(), btn),
		),
	)
	w.Show()

	return a
}

func textToGridRow(msg string) widget.TextGridRow {
	var cells []widget.TextGridCell
	for _, r := range []rune(msg) {
		cells = append(cells, widget.TextGridCell{Rune: r})
	}
	return widget.TextGridRow{
		Cells: cells,
	}
}
