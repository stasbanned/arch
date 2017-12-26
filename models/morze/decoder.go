package morze

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

//const dictFilename = "dictionary/eng.json"

func Json_reader() {
	rawDataIn, err := ioutil.ReadFile(path.dictFilename)
	if err != nil {
		log.Fatal("Cannot load dictionary:", err)
	}

	err = json.Unmarshal([]byte(rawDataIn), &dict)
	if err != nil {
		log.Fatal("Invalid dictionary format:", err)
	}
}

func Decoder(buffer string) string {
	result, ok := dict[buffer]
	if !ok {
		return "Error"
	}
	return result
}
