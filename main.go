package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
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

// 5BAA61E4C9B93F3F0682250B6CF8331B7EE68FD8
func main() {
	runtime.GOMAXPROCS(1)

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

	if *dictionary != "nil" && (*hash != "nil" || *hashes != "nil") {
		if *hash != "nil" {
			lowerCaseHash := strings.ToLower(*hash)
			passwords := readAndSplitFile(dictionary)
			checkPassword(passwords, lowerCaseHash)
		}

		if *hashes != "nil" {
			hashedPasswords := readAndSplitFile(hashes)
			for _, password := range hashedPasswords {
				lowerCaseHash := strings.ToLower(password)
				passwords := readAndSplitFile(dictionary)
				checkPassword(passwords, lowerCaseHash)
			}

		}
	}

	if *dictionaries != "nil" && (*hash != "nil" || *hashes != "nil") {
		passwordDicts, err := ioutil.ReadDir(*dictionaries)
		if err != nil {
			log.Fatal(err)
		}

		for _, dict := range passwordDicts {
			fileName := dict.Name()
			filePath := *dictionaries + "/" + fileName
			passwords := readAndSplitFile(&filePath)
			if *hash != "nil" {
				lowerCaseHash := strings.ToLower(*hash)
				checkPassword(passwords, lowerCaseHash)
			}

			if *hashes != "nil" {
				hashedPasswords := readAndSplitFile(hashes)
				for _, password := range hashedPasswords {
					lowerCaseHash := strings.ToLower(password)
					if err != nil {
						fmt.Println(err)
					}
					checkPassword(passwords, lowerCaseHash)
				}
			}
		}
	}

	os.Exit(0)
}
