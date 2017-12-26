package src

import (
	"github.com/andlabs/ui"
	"time"
	"../../presenters"
)

var (
	ChooseFileBtn      *ui.Button
	DecodeBtn          *ui.Button
	FileOutputTextView *ui.Label
	HandOutputTextView *ui.Label
	HandEditText       *ui.Entry
	Window             *ui.Window
	HandOutputText     = "Here will be the decoded text."
	FileOutputText     = "Here will be the decoded text."

	mainBox *ui.Box
)

func Show() {
	err := ui.Main(func() {
		mainBox = ui.NewHorizontalBox()
		mainBox.SetPadded(true)
		mainBox.Append(makeHandInputBlock(), true)
		mainBox.Append(makeFromFileBlock(), true)

		Window = ui.NewWindow("Morse decoder", 700, 300, false)
		Window.SetChild(mainBox)
		Window.OnClosing(func(*ui.Window) bool {
			ui.Quit()
			return true
		})
		Window.Show()
		go SetTextIntoOutputs()
	})
	if err != nil {
		panic(err)
	}
}

func SetTextIntoOutputs() {
	for  {
		ui.QueueMain(func() {
			HandOutputTextView.SetText(HandOutputText)
			FileOutputTextView.SetText(FileOutputText)
		})
		time.Sleep(time.Millisecond * 200)
	}
}

func makeHandInputBlock() *ui.Group {
	help := ui.NewLabel("Please, enter the morse code in field below")
	HandEditText = ui.NewEntry()
	HandOutputTextView = ui.NewLabel("Here will be the decoded text.")
	DecodeBtn = ui.NewButton("Decode")
	DecodeBtn.OnClicked(func(*ui.Button) {
		HandOutputText = presenters.OnDecodeClick(HandEditText.Text())
	})

	leftBox := ui.NewVerticalBox()
	leftBox.SetPadded(true)
	leftBox.Append(help, false)
	leftBox.Append(HandEditText, false)
	leftBox.Append(DecodeBtn, false)
	leftBox.Append(ui.NewHorizontalSeparator(), false)
	leftBox.Append(HandOutputTextView, true)

	leftGroup := ui.NewGroup("Hand HandEditText")
	leftGroup.SetMargined(true)
	leftGroup.SetChild(leftBox)
	return leftGroup
}

func makeFromFileBlock() *ui.Group {
	help := ui.NewLabel("Please, choose the .txt file with morse code")
	FileOutputTextView = ui.NewLabel("Here will be the decoded text.")
	ChooseFileBtn = ui.NewButton("Choose file")
	ChooseFileBtn.OnClicked(func(*ui.Button) {
		FileOutputText = presenters.OnChooseFileClick(ui.OpenFile(Window))
	})

	rightBox := ui.NewVerticalBox()
	rightBox.SetPadded(true)
	rightBox.Append(help, false)
	rightBox.Append(ChooseFileBtn, false)
	rightBox.Append(ui.NewHorizontalSeparator(), false)
	rightBox.Append(FileOutputTextView, true)

	rightGroup := ui.NewGroup("Open from file")
	rightGroup.SetMargined(true)
	rightGroup.SetChild(rightBox)
	return rightGroup
}
