package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"runtime"
)

func printUsage() {
	flag.Usage()
	fmt.Println("Example: crackr -h my-hash -f my-dictionary.txt")
	fmt.Println("Example 2: crackr -hf my-file-of-hashes.txt -d my-directory-of-dictionaries")
	os.Exit(1)
}

const resultsFile = "found_passwords"

// Performance tools & methods https://github.com/golang/go/wiki/Performance
func main() {
	runtime.GOMAXPROCS(1)
	// NOTE: https://markhneedham.com/blog/2017/01/31/go-multi-threaded-writing-csv-file/
	// Create results CSV
	file, err := os.Create(resultsFile + ".csv")
	checkError("Cannot create file", err)
	headers := [][]string{{"plaintext", "ciphertext", "hashing_algorithm"}}
	csv.NewWriter(file).WriteAll(headers)
	file.Close()

	hash := flag.String("h", "nil", "This is the hash of the password")
	hashes := flag.String("hf", "nil", "This is a file that contains multiple hashes to crack")
	dictionary := flag.String("f", "nil", "A single dictionary file with passwords to test")
	dictionaries := flag.String("d", "nil", "A directory with dictionary files")

	flag.Parse()

	if flag.NFlag() == 0 {
		printUsage()
	}

	if *hash == "nil" && *hashes == "nil" {
		panic("A hash is required to use crackr")
	}

	if *hash != "nil" && *hashes != "nil" {
		panic("Only one type of hash can be used!")
	}

	if *dictionary == "nil" && *dictionaries == "nil" {
		panic("A dictionary is required to use crackr")
	}

	if *dictionary != "nil" && *dictionaries != "nil" {
		panic("Only one type of dictionary can be used!")
	}
	// Handle combintation of single dictionary file and either a single hash or a hash file
	attackUsingSingleDictionary(dictionary, hash, hashes)
	// Handle combintation of a directory of dictionaries
	attackWithMultipleDictionaries(dictionaries, hash, hashes)

	os.Exit(0)
}
