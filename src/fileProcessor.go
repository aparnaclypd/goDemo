package main

import (
	"bufio"
	"fmt"
	"image"
	"log"
	"os"
	"strconv"
	"strings"

	"./helper"

	"github.com/disintegration/imaging"
)

type imageProcessorModel struct {
	srcImgName, destImgName string
	effects                 []string
}

func readInputFile(fileName string) {
	file, err := os.Open(fileName)

	if err != nil {
		log.Fatal(err)
		panic("Could not open input file.")
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var hasFinished = make(chan bool)
	var processedAllFiles = true

	for scanner.Scan() {

		newImageToBeProcessed := setDefaultValuesForModel(scanner.Text())

		scanner.Scan()

		for strings.TrimSpace(string(scanner.Text())) != "" {
			newImageToBeProcessed.effects = append(newImageToBeProcessed.effects, scanner.Text())
			scanner.Scan()
		}

		go processEffects(newImageToBeProcessed, hasFinished)
		processedAllFiles = processedAllFiles && <-hasFinished
	}

	if processedAllFiles {
		fmt.Println("Successfuly processed all images!")
	} else {
		fmt.Println("Could not process all images!")
	}

}

func processEffects(newImageToBeProcessed imageProcessorModel, hasFinished chan bool) {

	srcImg, er := imaging.Open("./images/" + newImageToBeProcessed.srcImgName)
	logHelper.LogError(er, "Could not open file")

	//Create output image in standard size
	destImg := standardizeImage(srcImg)

	//Start scanning effects and their values from file
	for _, line := range newImageToBeProcessed.effects {

		effectNameValue := strings.Split(line, " ")

		var value float64

		if len(effectNameValue) > 1 {
			value, _ = strconv.ParseFloat(effectNameValue[1], 64)
		}

		destImg = applyEffect(effectNameValue[0], value, destImg)

	}

	err := imaging.Save(destImg, "./images/"+newImageToBeProcessed.destImgName)
	logHelper.LogError(err, "Couldn't create output file.")
	fmt.Println("Created file : " + newImageToBeProcessed.destImgName)
	hasFinished <- true
}

func standardizeImage(srcImg image.Image) image.Image {
	destImg := imaging.CropAnchor(srcImg, 350, 350, imaging.Center)
	return imaging.Resize(destImg, 256, 0, imaging.Lanczos)
}

func setDefaultValuesForModel(scannedText string) imageProcessorModel {
	newImageToBeProcessed := imageProcessorModel{}
	fileNames := strings.Split(scannedText, " ")
	newImageToBeProcessed.srcImgName = fileNames[0]
	newImageToBeProcessed.destImgName = fileNames[1]
	return newImageToBeProcessed
}
