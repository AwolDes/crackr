package main

import (
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

	if *dictionary == "nil" && *dictionaries == "nil" {
		panic("A dictionary is required to use crackr")
	}

	lowerCaseHash := strings.ToLower(*hash)

	var passwords = []string{}

	if *dictionary != "nil" {
		passwordDict, err := ioutil.ReadFile(*dictionary)
		if err != nil {
			fmt.Print(err)
		}

		passwordString := string(passwordDict) // convert content to a 'string'
		passwords = strings.Split(passwordString, "\n")
	}

	checkPassword(passwords, lowerCaseHash)

	os.Exit(0)
}
