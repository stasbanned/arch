package tests

import (
	"testing"
	"../models/morze"
)

type testInterpreterPair struct {
	input  string
	output string
}


func BenchmarkTest(b *testing.B) {
	morze.SetDictFileName("../dictionary/eng.json")
	for i := 0; i < b.N; i++ {
		morze.SetInputFileName("../cash/hello.txt")
		morze.Interpreter()
    }
}


func BenchmarkTestParallel(b *testing.B) {
	morze.SetDictFileName("../dictionary/eng.json")
	for i := 0; i < b.N; i++ {
		morze.SetInputFileName("../cash/hello.txt")
		morze.Threader()
    }
}