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
	w := a.NewWindow("counterbaal")
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
	overridenumber.PlaceHolder = "'73'"

	numberButton := widget.NewButton("queue manual number", func() {
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

// Licenes of dependencies:

// 1. clipboard

// MIT License

// Copyright (c) 2021 Changkun Ou <contact@changkun.de>

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:

// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

// 2. Fyne

// BSD 3-Clause License

// Copyright (C) 2018 Fyne.io developers (see AUTHORS)
// All rights reserved.

// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are met:
//     * Redistributions of source code must retain the above copyright
//       notice, this list of conditions and the following disclaimer.
//     * Redistributions in binary form must reproduce the above copyright
//       notice, this list of conditions and the following disclaimer in the
//       documentation and/or other materials provided with the distribution.
//     * Neither the name of Fyne.io nor the names of its contributors may be
//       used to endorse or promote products derived from this software without
//       specific prior written permission.

// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND
// ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED
// WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
// DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER BE LIABLE FOR ANY
// DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES
// (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES;
// LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND
// ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
// (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS
// SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
