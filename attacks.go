package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func attackUsingSingleDictionary(dictionary *string, hash *string, hashes *string) {
	var foundPasswords []string
	if *dictionary != "nil" && (*hash != "nil" || *hashes != "nil") {
		if *hash != "nil" {
			lowerCaseHash := strings.ToLower(*hash)
			passwords := readAndSplitFile(dictionary)
			checkPassword(passwords, &foundPasswords, lowerCaseHash)
		}

		if *hashes != "nil" {
			hashedPasswords := readAndSplitFile(hashes)
			for _, password := range hashedPasswords {
				lowerCaseHash := strings.ToLower(password)
				passwords := readAndSplitFile(dictionary)
				checkPassword(passwords, &foundPasswords, lowerCaseHash)
			}

		}
	}
}

func attackWithMultipleDictionaries(dictionaries *string, hash *string, hashes *string) {
	var foundPasswords []string
	if *dictionaries != "nil" && (*hash != "nil" || *hashes != "nil") {
		passwordDicts, err := ioutil.ReadDir(*dictionaries)
		checkError("Could not read directory: ", err)

		for _, dict := range passwordDicts {
			fileName := dict.Name()
			filePath := *dictionaries + "/" + fileName
			passwords := readAndSplitFile(&filePath)
			if *hash != "nil" {
				lowerCaseHash := strings.ToLower(*hash)
				checkPassword(passwords, &foundPasswords, lowerCaseHash)
			}

			if *hashes != "nil" {
				hashedPasswords := readAndSplitFile(hashes)
				for _, password := range hashedPasswords {
					lowerCaseHash := strings.ToLower(password)
					if err != nil {
						fmt.Println(err)
					}
					checkPassword(passwords, &foundPasswords, lowerCaseHash)
				}
			}
		}
	}
}
