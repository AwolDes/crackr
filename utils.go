package main

import (
	"encoding/hex"
	"fmt"
	"hash"
	"io/ioutil"
	"log"
	"strings"
)

func readAndSplitFile(file *string) []string {
	fileBytes, err := ioutil.ReadFile(*file)
	if err != nil {
		fmt.Println(err)
	}
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
