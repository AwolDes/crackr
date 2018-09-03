package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"strings"

	"golang.org/x/crypto/sha3"
)

func readAndSplitFile(file *string) []string {
	fileBytes, err := ioutil.ReadFile(*file)
	if err != nil {
		fmt.Println(err)
	}
	fileString := string(fileBytes)
	stringArray := strings.Split(fileString, "\n")
	return stringArray
}

func getHash(hashType string, plaintext string) string {
	switch hashingAlgo := hashType; hashingAlgo {
	case "sha1":
		hasher := sha1.New()
		hasher.Write([]byte(plaintext))
		sha1 := hex.EncodeToString(hasher.Sum(nil))
		return sha1
	case "md5":
		hasher := md5.New()
		hasher.Write([]byte(plaintext))
		md5 := hex.EncodeToString(hasher.Sum(nil))
		return md5
	case "sha256":
		hasher := sha256.New()
		hasher.Write([]byte(plaintext))
		sha256 := hex.EncodeToString(hasher.Sum(nil))
		return sha256
	case "sha512":
		hasher := sha512.New()
		hasher.Write([]byte(plaintext))
		sha512 := hex.EncodeToString(hasher.Sum(nil))
		return sha512
	case "sha3_256":
		hasher := sha3.New256()
		hasher.Write([]byte(plaintext))
		sha3_256 := hex.EncodeToString(hasher.Sum(nil))
		return sha3_256
	case "sha3_512":
		hasher := sha3.New512()
		hasher.Write([]byte(plaintext))
		sha3_512 := hex.EncodeToString(hasher.Sum(nil))
		return sha3_512
	default:
		panic("Hash type not supported!")
	}
}

func checkPassword(passwords []string, hash string) {
	for _, password := range passwords {
		attempts := 0
		hashAlgorithims := []string{"sha1", "md5", "sha256", "sha512", "sha3_256", "sha3_512"}
		for _, hashAlgorithim := range hashAlgorithims {
			hashedPassword := getHash(hashAlgorithim, password)
			if hashedPassword == hash {
				fmt.Println("Match!")
				fmt.Println(password, hashAlgorithim)
			}
			attempts++
		}
	}
}
