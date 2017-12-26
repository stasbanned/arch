package tests

import (
	"testing"
	"../models/morze"
)

type testDecoderPair struct {
	input  string
	output string
}

var decoderTests = []testDecoderPair{
	{"10002","b"},
	{" ","Error"},
	{"011 ","w"},
	{"0112","w"},
	{"1010 ","c"},
}

func TestDecoder(t *testing.T) {
	for _, pair := range decoderTests {
		morze.SetDictFileName("../dictionary/eng.json")
		morze.SetInputFileName("../cash/input.txt")
		v := morze.Decoder(pair.input)
		if v != pair.output {
			t.Error(
				"For", pair.input,
				"expected", pair.output,
				"got", v,
			)
		}
	}
}