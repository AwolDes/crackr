package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"runtime"
	"strings"
)

func printUsage() {
	flag.Usage()
	fmt.Println("Example: crackr -h my-hash -f my-dictionary.txt")
	fmt.Println("Example 2: crackr -hf my-file-of-hashes.txt -d my-directory-of-dictionaries")
	os.Exit(1)
}

const ResultsFile = "found_passwords"

func main() {
	runtime.GOMAXPROCS(1)
	// NOTE: https://markhneedham.com/blog/2017/01/31/go-multi-threaded-writing-csv-file/
	// Create results CSV
	file, err := os.Create(ResultsFile + ".csv")
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

	var foundPasswords []string
	// Handle combintation of single dictionary file and either a single hash or a hash file
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
	// Handle combintation of a directory of dictionaries
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

	os.Exit(0)
}
