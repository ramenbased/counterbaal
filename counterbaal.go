package main

import (
	"strconv"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"

	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"golang.design/x/clipboard"
)

// go build -ldflags "-w -s -H=windowsgui"

var number int64
var numberb4 int64
var history []string

func main() {

	start := time.Now()

	a := app.New()
	w := a.NewWindow("clipbaal")
	w.Resize(fyne.NewSize(250, 600))

	// ---- Timer

	data := binding.BindStringList(
		&[]string{},
	)

	list := widget.NewListWithData(data,
		func() fyne.CanvasObject {
			return widget.NewLabel("template")
		},
		func(i binding.DataItem, o fyne.CanvasObject) {
			o.(*widget.Label).Bind(i.(binding.String))
		})

	// ---- Widgets

	nexttext := widget.NewLabel("'mybaalrun-01'")
	in := widget.NewEntry()
	in.PlaceHolder = "'mybaalrun'"

	ngButton := widget.NewButton("ng", func() {
		//timesetup
		t := time.Now()
		elapsed := t.Sub(start)

		//
		nexttext.SetText(counter(in.Text, "next"))
		var val string

		// -6 means no skip ng number
		if numberb4 == -6 {
			val = cformat(number-1, in.Text)
		} else {
			val = cformat(numberb4, in.Text)
		}
		val_last := historyadd(val)

		// kick out first run
		if len(history) < 2 {
			data.Append(val_last)
		} else {
			data.Append(val_last + "\t\t" + strconv.FormatFloat(elapsed.Seconds(), 'f', 2, 64) + " sec")
		}

		//misc
		numberb4 = -6
		start = time.Now()
		list.ScrollToBottom()
	})

	// backButton := widget.NewButton("back", func() {
	// 	nexttext.SetText(counter(in.Text, "back"))
	// 	start = time.Now()
	// })

	resetButton := widget.NewButton("reset", func() {
		nexttext.SetText(counter(in.Text, "reset"))
		data.Set(nil)
		start = time.Now()
	})

	overridenumber := widget.NewEntry()
	overridenumber.PlaceHolder = "enter manual No. '3'"

	numberButton := widget.NewButton("skip to No. in ng", func() {
		i, _ := strconv.ParseInt(overridenumber.Text, 10, 0)
		numberb4 = number
		number = i - 1
	})

	// ---- Lay

	topbox := container.NewVBox(
		in,
		nexttext,
		ngButton,
		//backButton,
		//resetButton,
		overridenumber,
		numberButton,
	)

	grid := fyne.NewContainerWithLayout(layout.NewBorderLayout(topbox, resetButton, nil, nil),
		topbox,
		list,
		resetButton,
	)

	// ---- Render

	w.SetContent(grid)
	w.ShowAndRun()
}

// ---- funcs

func historyadd(a string) string {
	history = append(history, a)
	if len(history) < 2 {
		return "simply ctrl+v into d2 | History:"
	} else {
		return history[len(history)-1]
	}
}

func counter(name string, action string) string {
	var rv string
	switch a := action; a {
	case "next":
		number += 1
		rv = cformat(number, name)
	case "back":
		if number > 0 {
			number -= 1
		} else {
			number = 0
		}
		rv = cformat(number, name)
	case "reset":
		number = 0
		history = nil
		rv = cformat(number, name)
	}
	clipboard.Write(clipboard.FmtText, []byte(rv))
	return rv
}

func cformat(i int64, name string) string {
	if i < 10 {
		return name + "-0" + strconv.FormatInt(i, 10)
	} else {
		return name + "-" + strconv.FormatInt(i, 10)
	}
}
