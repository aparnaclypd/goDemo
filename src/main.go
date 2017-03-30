package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	inputDir := "./input/"
	hasFinished := make(chan bool)
	exitProgram := true

	files, _ := ioutil.ReadDir(inputDir)
	for _, file := range files {
		go readFile(inputDir+file.Name(), hasFinished)
		exitProgram = exitProgram && <-hasFinished
	}

	if exitProgram {
		fmt.Println("----Done processing files!----")
	}

}
