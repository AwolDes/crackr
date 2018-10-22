package main

import (
	"sync"
)

var hashAlgorithmFuncs = map[string]func(string) string{
	"sha1":     sha1Hash,
	"md5":      md5Hash,
	"sha256":   sha256Hash,
	"sha512":   sha512Hash,
	"sha3_256": sha3_256Hash,
	"sha3_512": sha3_512Hash,
}

var hashAlgorithmOptions = []string{"sha1", "md5", "sha256", "sha512", "sha3_256", "sha3_512"}

/*
	Returns the hashed version of the plaintext
*/
func getHash(hashType string, plaintext string) string {
	// note: a map function is faster than a switch
	// source: https://hashrocket.com/blog/posts/switch-vs-map-which-is-the-better-way-to-branch-in-go
	hashAlgorithm := hashAlgorithmFuncs[hashType]
	return hashAlgorithm(plaintext)
}

/*
	Check if a password has already been found
*/
func checkFoundPasswords(foundPasswords *PasswordsFound, hashedPassword string) bool {
	// this is only thread safe when foundPasswords is copied, and not a pointer
	// for _, foundPassword := range foundPasswords.passwords {
	// 	if foundPassword == hashedPassword {
	// 		return true
	// 	}
	// }
	foundPasswords.appendPassword(hashedPassword)
	return false
}

/*
	If a password is found, write it to the CSV to be analysed later.
*/
func foundPassword(password string, hashedPassword string, hashAlgorithim string, foundPasswords *PasswordsFound) {
	csvWriter.writeChanges([]string{password, hashedPassword, hashAlgorithim})
}

/*
	For each hash algorithm, hash each password and see if any match the given hash
*/
func checkPassword(passwords []string, hash string) {
	// for each hash algorithm, check all passwords
	var wg sync.WaitGroup
	for _, hashAlgorithim := range hashAlgorithmOptions {
		wg.Add(1)
		go func(hashAlgorithim string, passwords []string, hash string) {
			defer wg.Done()
			for _, password := range passwords {
				hashedPassword := getHash(hashAlgorithim, password)
				if hashedPassword == hash && foundPasswords.appendPassword(hashedPassword) == false {
					csvWriter.writeChanges([]string{password, hashedPassword, hashAlgorithim})
				}
			}
		}(hashAlgorithim, passwords, hash)
	}
	wg.Wait()
}
