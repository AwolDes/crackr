package main

import (
	"io/ioutil"
	"strings"
	"sync"
)

/*
	This function chunks a password dictionary into 4 chunks using
	the util function chunkArray
*/
func chunkPasswordDictionary(dictionary *string) [][]string {
	passwords := readAndSplitFile(dictionary)
	return chunkArray(passwords, 4)
}

/*
	This function starts a goroutine for each password, and checks if that password is in the
	given password dictionary chunk
*/
func checkPasswords(dictionaryChunk []string, foundPasswords []string, hashedPasswords []string) {
	// var wg sync.WaitGroup
	// foundPasswordsChannel := make(chan []string, 1)
	// defer close(foundPasswordsChannel)
	for _, password := range hashedPasswords {
		lowerCasePassword := strings.ToLower(password)
		checkPassword(dictionaryChunk, &foundPasswords, lowerCasePassword)
		// wg.Add(1)
		// go func(dictionaryChunk []string, foundPasswords []string, password string) {
		// 	defer wg.Done()
		// }(dictionaryChunk, foundPasswords, password)
	}
	// wg.Wait()
}

/*
	This function starts a goroutine for all dictionary chunks, and then calls
	checkPasswords to see if any passwords are in any dictionary chunks
*/
func searchChunkedDictionary(chunkedDictionary [][]string, hashedPasswords []string, foundPasswords []string) {
	var wg sync.WaitGroup
	for _, passwordChunk := range chunkedDictionary {
		wg.Add(1)
		go func(passwordChunk []string, foundPasswords []string) {
			defer wg.Done()
			checkPasswords(passwordChunk, foundPasswords, hashedPasswords)
		}(passwordChunk, foundPasswords)
	}
	// Does not sync - race condition?
	wg.Wait()

}

/*
	This function handles logic for a dictionary that is just a single file
*/
func attackUsingSingleDictionary(dictionary *string, hash *string, hashes *string) {
	var foundPasswords []string
	if *dictionary != "nil" && (*hash != "nil" || *hashes != "nil") {
		if *hash != "nil" {
			lowerCaseHash := strings.ToLower(*hash)
			chunkedDictionary := chunkPasswordDictionary(dictionary)
			searchChunkedDictionary(chunkedDictionary, []string{lowerCaseHash}, foundPasswords)
		}

		if *hashes != "nil" {
			hashedPasswords := readAndSplitFile(hashes)
			chunkedDictionary := chunkPasswordDictionary(dictionary)
			searchChunkedDictionary(chunkedDictionary, hashedPasswords, foundPasswords)
		}
	}
}

/*
	This function handles logic for a directory of dictionaries
*/
func attackWithMultipleDictionaries(dictionaries *string, hash *string, hashes *string) {
	var foundPasswords []string
	if *dictionaries != "nil" && (*hash != "nil" || *hashes != "nil") {
		passwordDicts, err := ioutil.ReadDir(*dictionaries)
		checkError("Could not read directory: ", err)

		// Interesting, paralellising this loop decreases performance
		for _, dictionary := range passwordDicts {
			fileName := dictionary.Name()
			filePath := *dictionaries + "/" + fileName
			if *hash != "nil" {
				lowerCaseHash := strings.ToLower(*hash)
				chunkedDictionary := chunkPasswordDictionary(&filePath)
				searchChunkedDictionary(chunkedDictionary, []string{lowerCaseHash}, foundPasswords)
			}

			if *hashes != "nil" {
				hashedPasswords := readAndSplitFile(hashes)
				chunkedDictionary := chunkPasswordDictionary(&filePath)
				searchChunkedDictionary(chunkedDictionary, hashedPasswords, foundPasswords)
			}
		}
	}
}
