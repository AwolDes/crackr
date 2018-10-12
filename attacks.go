package main

import (
	"io/ioutil"
	"strings"
	"sync"
)

func chunkPasswordDictionary(dictionary *string) [][]string {
	passwords := readAndSplitFile(dictionary)
	return chunkArray(passwords, 4)
}

func checkPasswords(passwordChunk []string, foundPasswords []string, hashedPasswords []string) {
	var wg sync.WaitGroup
	for _, password := range hashedPasswords {
		wg.Add(1)
		go func(passwordChunk []string, foundPasswords []string, password string) {
			defer wg.Done()
			lowerCasePassword := strings.ToLower(password)
			checkPassword(passwordChunk, &foundPasswords, lowerCasePassword)
		}(passwordChunk, foundPasswords, password)
	}
}

func searchChunkedDictionary(chunkedDictionary [][]string, hashedPasswords []string, foundPasswords []string) {
	var wg sync.WaitGroup
	for _, passwordChunk := range chunkedDictionary {
		wg.Add(1)
		go func(passwordChunk []string, foundPasswords []string) {
			defer wg.Done()
			checkPasswords(passwordChunk, foundPasswords, hashedPasswords)
		}(passwordChunk, foundPasswords)
	}
	wg.Wait()
}

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
			chunkedDictionary := chunkPasswordDictionary(dictionary)
			searchChunkedDictionary(chunkedDictionary, hashedPasswords, foundPasswords)
		}
	}
}

func attackWithMultipleDictionaries(dictionaries *string, hash *string, hashes *string) {
	var foundPasswords []string
	if *dictionaries != "nil" && (*hash != "nil" || *hashes != "nil") {
		passwordDicts, err := ioutil.ReadDir(*dictionaries)
		checkError("Could not read directory: ", err)

		for _, dictionary := range passwordDicts {
			fileName := dictionary.Name()
			filePath := *dictionaries + "/" + fileName
			passwordDictionary := readAndSplitFile(&filePath)
			if *hash != "nil" {
				lowerCaseHash := strings.ToLower(*hash)
				checkPassword(passwordDictionary, &foundPasswords, lowerCaseHash)
			}

			if *hashes != "nil" {
				hashedPasswords := readAndSplitFile(hashes)
				chunkedDictionary := chunkPasswordDictionary(&filePath)
				searchChunkedDictionary(chunkedDictionary, hashedPasswords, foundPasswords)
			}
		}
	}
}
