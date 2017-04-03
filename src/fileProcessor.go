package main

import (
	"bufio"
	"fmt"
	"image"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"

	"github.com/disintegration/imaging"
)

type imageProcessorModel struct {
	srcImgName, destImgName string
	effects                 []string
}

func readInputFile(fileName string) error {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
		return NewCouldNotOpenFileError("Could not open file!")
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	var wg sync.WaitGroup
	var processNextLine = true

	for processNextLine {

		processNextLine = scanner.Scan()
		if !processNextLine {
			err := scanner.Err()
			if err == nil {
				return NewScanInputFileError("No data present in input file")
			}
			return NewScanInputFileError(err.Error())
		}

		newImageToBeProcessed := setDefaultValuesForModel(scanner.Text())

		processNextLine = scanner.Scan()
		if !processNextLine {
			err := scanner.Err()
			if err == nil {
				return NewScanInputFileError(fmt.Sprintf("No Effects provided for %s - %s combo.", newImageToBeProcessed.srcImgName, newImageToBeProcessed.destImgName))
			}
			return NewScanInputFileError(err.Error())
		}

		for strings.TrimSpace(string(scanner.Text())) != "" {
			newImageToBeProcessed.effects = append(newImageToBeProcessed.effects, scanner.Text())

			processNextLine = scanner.Scan()
			if !processNextLine {
				err := scanner.Err()
				if err == nil {
					break
				}
				return NewScanInputFileError(err.Error())
			}
		}
		wg.Add(1)
		go processEffects(newImageToBeProcessed, &wg)
	}

	wg.Wait()
	return nil
}

func processEffects(newImageToBeProcessed imageProcessorModel, wg *sync.WaitGroup) {

	srcImg, err := imaging.Open("./images/" + newImageToBeProcessed.srcImgName)
	if err != nil {
		log.Fatal(err)
		//msg := fmt.Sprintf("Could not open image: %s.", newImageToBeProcessed.srcImgName)
		//return NewCouldNotOpenFileError(msg)
	}

	//Create output image in standard size
	destImg := standardizeImage(srcImg)

	//Start scanning effects and their values from file
	for _, line := range newImageToBeProcessed.effects {

		effectNameValue := strings.Split(line, " ")

		var value float64

		if len(effectNameValue) > 1 {
			value, err = strconv.ParseFloat(effectNameValue[1], 64)
			if err != nil {
				//return NewScanValueError(fmt.Sprintf("Invalid input value : %s", effectNameValue[1]))
			}
		}

		destImg = applyEffect(effectNameValue[0], value, destImg)

	}

	err = imaging.Save(destImg, "./images/"+newImageToBeProcessed.destImgName)
	if err != nil {
		log.Fatal(err)
		//		return NewCouldNotOpenFileError("Could not save output file: %s. Errors: %s", newImageToBeProcessed.destImgName, err.Error())
	}

	fmt.Println("Created file : " + newImageToBeProcessed.destImgName)
	wg.Done()
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
