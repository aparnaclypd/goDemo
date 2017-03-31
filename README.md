# goDemo
This project is about adding filters to the image and creting output image after application of those filters.

Input is taken in form of txt files. 
Format of txt files is as follows:
Line 1: sourceImage.png outputImageName.jpeg
Line 2: filterName <intensity>
Line n: filterName <intensity>

A sample input file will look like:

nature.png editedNature.png
blur 2
grayscale
sharpness 30

Currently, in the code we take multiple input files and process them to create output files.
Processing of each file triggers a goroutine.

Folder hierarchy:

"input" folder contains all the txt files which are to be processed.
"images" folder contains source and output images
