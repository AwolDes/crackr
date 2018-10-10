package main

import (
	"encoding/hex"
	"hash"
	"io/ioutil"
	"log"
	"strings"
)

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
