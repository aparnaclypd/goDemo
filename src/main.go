package main

import "fmt"

func main() {

	err := readInputFile("./input/input.txt")

	switch err.(type) {

	case *CouldNotOpenFileError:
		fmt.Println(err.Error())
	case *ScanInputFileError:
		fmt.Println(err.Error())
	case *ScanValueError:
		fmt.Println(err.Error())

	}

}
