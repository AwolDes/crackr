package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"

	"golang.org/x/crypto/sha3"
)

func getHash(hashType string, plaintext string) string {
	switch hashingAlgo := hashType; hashingAlgo {
	case "sha1":
		hasher := sha1.New()
		sha1 := hashText(hasher, plaintext)
		return sha1
	case "md5":
		hasher := md5.New()
		md5 := hashText(hasher, plaintext)
		return md5
	case "sha256":
		hasher := sha256.New()
		sha256 := hashText(hasher, plaintext)
		return sha256
	case "sha512":
		hasher := sha512.New()
		sha512 := hashText(hasher, plaintext)
		return sha512
	case "sha3_256":
		hasher := sha3.New256()
		sha3_256 := hashText(hasher, plaintext)
		return sha3_256
	case "sha3_512":
		hasher := sha3.New512()
		sha3_512 := hashText(hasher, plaintext)
		return sha3_512
	default:
		panic("Hash type not supported!")
	}
}

func checkFoundPasswords(foundPasswords *[]string, hashedPassword string) bool {
	for _, foundPassword := range *foundPasswords {
		if foundPassword == hashedPassword {
			return true
		}
	}
	return false
}

func checkPassword(passwords []string, foundPasswords *[]string, hash string) {
	// for each password, check all hash algorithms
	for _, password := range passwords {
		hashAlgorithims := []string{"sha1", "md5", "sha256", "sha512", "sha3_256", "sha3_512"}
		for _, hashAlgorithim := range hashAlgorithims {
			hashedPassword := getHash(hashAlgorithim, password)
			if len(*foundPasswords) > 0 {
				// check if a password has already been found
				if !checkFoundPasswords(foundPasswords, hashedPassword) {
					if hashedPassword == hash {
						fmt.Println("Matched!")
						fmt.Println(password, hashAlgorithim)
						writeCSV(resultsFile, []string{password, hashedPassword, hashAlgorithim})
						*foundPasswords = append(*foundPasswords, hashedPassword)
					}
				}
			} else {
				if hashedPassword == hash {
					fmt.Println("Match!")
					fmt.Println(password, hashAlgorithim)
					writeCSV(resultsFile, []string{password, hashedPassword, hashAlgorithim})
					*foundPasswords = append(*foundPasswords, hashedPassword)
				}
			}

		}
	}
}
