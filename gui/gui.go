package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"log"
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
		err := download(urlInput.Text, outputInput.Text)
		if err != nil {
			log.Println(err)
		}
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
