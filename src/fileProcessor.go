package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/disintegration/imaging"

	"./helper"
)

func readFile(fileName string, hasFinished chan bool) {

	bytes, errr := ioutil.ReadFile(fileName)
	logHelper.LogError(errr, "Could not open input file!")

	linesInFile := strings.Split(string(bytes), "\n")

	processEffects(linesInFile)

	hasFinished <- true

}

func processEffects(linesInFile []string) {

	fileNames := strings.Split(linesInFile[0], " ")

	srcFileName := fileNames[0]
	destFileName := fileNames[1]

	srcImg, er := imaging.Open("./images/" + srcFileName)
	logHelper.LogError(er, "Could not open file")

	//Create output image in standard size
	destImg := imaging.CropAnchor(srcImg, 350, 350, imaging.Center)
	destImg = imaging.Resize(destImg, 256, 0, imaging.Lanczos)

	//Start scanning effects and their values from file
	for index := 1; strings.TrimSpace(string(linesInFile[index])) != ""; index++ {

		scannedTokens := strings.Split(linesInFile[index], " ")

		var effect Effect
		var value float64

		if len(scannedTokens) > 1 {
			value, _ = strconv.ParseFloat(scannedTokens[1], 64)
		}

		switch strings.ToLower(scannedTokens[0]) {

		case "blur":
			effect = blur{value}
			break

		case "rotate":
			effect = rotate{value}
			break

		case "contrast":
			effect = contrast{value}
			break

		case "grayscale":
			effect = grayscale{}
			break

		case "sharpness":
			effect = sharpness{value}
			break
		}

		//Apply effect
		destImg = effect.apply(destImg)
	}
	err := imaging.Save(destImg, "./images/"+destFileName)
	logHelper.LogError(err, "Couldn't create output file.")
	fmt.Println("Created file : " + destFileName)
}
