package main

import (
	"./models/morze"
	// "./views/src"
	"fmt"
)

func main() {
	morze.DefaultValues()
	fmt.Println(morze.Threader())
	// src.Show()
}
