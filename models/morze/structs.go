package morze

const threadsNumber = 10

var dict map[string]string;
var path = PathCode{}
var maps = [threadsNumber]string{}
var sem = make(chan int, 1)


type PathCode struct {
	dictFilename string
	inputFileName string
}

func addCommonMemory(str string, index int) {
	sem <- 1
	maps[index] = str
	<- sem
}

func SetDictFileName(dictFilename string) {
	path.dictFilename = dictFilename
}

func SetInputFileName(inputFileName string) {
	path.inputFileName = inputFileName
}


func DefaultValues() {
	path.dictFilename = "dictionary/eng.json"
	path.inputFileName = "cash/hello.txt"
	Json_reader()
}
