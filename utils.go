package main

import (
	"encoding/hex"
	"hash"
	"io/ioutil"
	"log"
	"strings"
)

var numCPU = 4

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

func readAndSplitFile(file *string) []string {
	fileBytes, err := ioutil.ReadFile(*file)
	checkError("Could not split file: ", err)
	fileString := string(fileBytes)
	stringArray := strings.Split(fileString, "\n")
	return stringArray
}

func hashText(hasher hash.Hash, plaintext string) string {
	hasher.Write([]byte(plaintext))
	cipherText := hex.EncodeToString(hasher.Sum(nil))
	return cipherText
}

func checkError(message string, err error) {
	if err != nil {
		log.Fatal(message, err)
	}
}
