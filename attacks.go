package main

import (
	"io/ioutil"
	"os"
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
func checkPasswords(dictionaryChunk []string, hashedPasswords []string) {
	var wg sync.WaitGroup
	for _, password := range hashedPasswords {
		wg.Add(1)
		go func(dictionaryChunk []string, password string) {
			defer wg.Done()
			lowerCasePassword := strings.ToLower(password)
			checkPassword(dictionaryChunk, lowerCasePassword)
		}(dictionaryChunk, password)
	}
	wg.Wait()
}

/*
	This function starts a goroutine for all dictionary chunks, and then calls
	checkPasswords to see if any passwords are in any dictionary chunks
*/
func searchChunkedDictionary(chunkedDictionary [][]string, hashedPasswords []string) {
	var wg sync.WaitGroup
	for _, passwordChunk := range chunkedDictionary {
		wg.Add(1)
		go func(passwordChunk []string, hashedPasswords []string) {
			defer wg.Done()
			checkPasswords(passwordChunk, hashedPasswords)
		}(passwordChunk, hashedPasswords)
	}
	wg.Wait()

}

/*
	This function handles logic for a dictionary that is just a single file
*/
func attackUsingSingleDictionary(dictionary *string, hash *string, hashes *string) {
	if *dictionary != "nil" && (*hash != "nil" || *hashes != "nil") {

		// foundPasswordsChan := make(chan PasswordsFound)
		if *hash != "nil" {
			lowerCaseHash := strings.ToLower(*hash)
			chunkedDictionary := chunkPasswordDictionary(dictionary)
			searchChunkedDictionary(chunkedDictionary, []string{lowerCaseHash})
		}

		if *hashes != "nil" {
			hashedPasswords := readAndSplitFile(hashes)
			chunkedDictionary := chunkPasswordDictionary(dictionary)
			searchChunkedDictionary(chunkedDictionary, hashedPasswords)
		}
	}
}

/*
	This function handles logic for a directory of dictionaries
*/
func attackWithMultipleDictionaries(dictionaries *string, hash *string, hashes *string) {
	if *dictionaries != "nil" && (*hash != "nil" || *hashes != "nil") {
		passwordDicts, err := ioutil.ReadDir(*dictionaries)
		checkError("Could not read directory: ", err)

		var wg sync.WaitGroup
		for _, dictionary := range passwordDicts {
			wg.Add(1)
			go func(dictionary os.FileInfo) {
				defer wg.Done()
				fileName := dictionary.Name()
				filePath := *dictionaries + "/" + fileName
				if *hash != "nil" {
					lowerCaseHash := strings.ToLower(*hash)
					chunkedDictionary := chunkPasswordDictionary(&filePath)
					searchChunkedDictionary(chunkedDictionary, []string{lowerCaseHash})
				}

				if *hashes != "nil" {
					hashedPasswords := readAndSplitFile(hashes)
					chunkedDictionary := chunkPasswordDictionary(&filePath)
					searchChunkedDictionary(chunkedDictionary, hashedPasswords)
				}
			}(dictionary)
		}
		wg.Wait()
	}
}
