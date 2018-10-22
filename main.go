package main

import (
	"flag"
	"fmt"
	"os"
	"runtime"
	"sync"
)

func printUsage() {
	flag.Usage()
	fmt.Println("Example: crackr -h my-hash -f my-dictionary.txt")
	fmt.Println("Example 2: crackr -hf my-file-of-hashes.txt -d my-directory-of-dictionaries")
	os.Exit(1)
}

const resultsFile = "found_passwords"

var csvWriter = newThreadSafeCsvWriter()

var foundPasswords = PasswordsFound{
	mutex:     &sync.Mutex{},
	passwords: []string{},
}

// Performance tools & methods https://github.com/golang/go/wiki/Performance
func main() {
	runtime.GOMAXPROCS(8)

	headers := []string{"plaintext", "ciphertext", "hashing_algorithm"}
	csvWriter.writeChanges(headers)

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
	csvWriter.flush()
	os.Exit(0)
}
