package gui

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

type Downloader interface {
}

func New(download func(string, string) error) fyne.App {
	a := app.New()
	w := a.NewWindow("m3u8 downloader")
	w.Resize(fyne.Size{Width: 500, Height: 200})

	urlInput := widget.NewEntry()
	outputInput := widget.NewEntry()

	logInfo := widget.NewTextGrid()
	logInfo.Resize(fyne.Size{Height: 100})

	btn := widget.NewButton("Download", nil)
	btn.Resize(fyne.Size{Width: 80})
	btn.OnTapped = func() {
		btn.Disable()
		defer btn.Enable()

		url := urlInput.Text
		file := outputInput.Text

		logInfo.SetRow(len(logInfo.Rows), textToGridRow(fmt.Sprintf("Start download %s", url)))
		err := download(url, file)
		if err != nil {
			logInfo.SetRow(len(logInfo.Rows)+1, textToGridRow(err.Error()))
		}
		logInfo.SetRow(len(logInfo.Rows)+1, textToGridRow(fmt.Sprintf("File saved %s", file)))
	}

	w.SetContent(
		container.New(layout.NewGridLayout(1),
			widget.NewForm(
				widget.NewFormItem("URL", urlInput),
				widget.NewFormItem("File", outputInput),
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
