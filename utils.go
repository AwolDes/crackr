package main

import (
	"io/ioutil"
	"log"
	"strings"
)

/*
	Given an array, this function will split it into the specified
	amount of chunks
*/
func chunkArray(array []string, chunks int) [][]string {
	arrayLength := len(array)
	chunkSize := (arrayLength + 1) / chunks
	chunkedArray := [][]string{}
	for i := 0; i < arrayLength; i += chunkSize {
		chunkEnd := i + chunkSize

		if chunkEnd > arrayLength {
			chunkEnd = arrayLength
		}

		chunkedArray = append(chunkedArray, array[i:chunkEnd])
	}
	return chunkedArray
}

/*
	Splits a file into an array
*/
func readAndSplitFile(file *string) []string {
	fileBytes, err := ioutil.ReadFile(*file)
	checkError("Could not split file: ", err)
	fileString := string(fileBytes)
	stringArray := strings.Split(fileString, "\n")
	return stringArray
}

/*
	Generic error handling as go requires explicit handling of errors
*/
func checkError(message string, err error) {
	if err != nil {
		log.Fatal(message, err)
	}
}
