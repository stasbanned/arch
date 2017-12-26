package presenters

import (
	"../models/morze"
)

func OnDecodeClick(s string) string {
	morze.SetUserInput(s)
	return morze.Interpreter()
}

func OnChooseFileClick(s string) string {
	morze.SetInputFileName(s)
	return morze.Interpreter()
}
