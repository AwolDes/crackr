package main

var hashAlgorithms = map[string]func(string) string{
	"sha1":     sha1Hash,
	"md5":      md5Hash,
	"sha256":   sha256Hash,
	"sha512":   sha512Hash,
	"sha3_256": sha3_256Hash,
	"sha3_512": sha3_512Hash,
}

/*
	Gets a hash for a given string and hash algorithm
*/
func getHash(hashType string, plaintext string) string {
	// note: a map function is faster than a switch
	// source: https://hashrocket.com/blog/posts/switch-vs-map-which-is-the-better-way-to-branch-in-go
	for _, hashAlgorithm := range hashAlgorithms {
		return hashAlgorithm(plaintext)
	}
	panic("Hash type not supported!")
}

/*
	Check if a password has already been found
*/
func checkFoundPasswords(foundPasswords *[]string, hashedPassword string) bool {
	for _, foundPassword := range *foundPasswords {
		if foundPassword == hashedPassword {
			return true
		}
	}
	return false
}

/*
	If a password is found, write it to the CSV to be analysed later.
*/
func foundPassword(password string, hashedPassword string, hashAlgorithim string, foundPasswords *[]string, foundPasswordsChannel chan []string) {
	csvWriter.writeChanges([]string{password, hashedPassword, hashAlgorithim})
	*foundPasswords = append(*foundPasswords, hashedPassword)
	foundPasswordsChannel <- *foundPasswords
}

/*
	For each hash algorithm, hash each password and see if any match the given hash
*/
func checkPassword(passwords []string, foundPasswords *[]string, hash string, foundPasswordsChannel chan []string) {
	// for each hash algorithm, check all passwords
	hashAlgorithims := []string{"sha1", "md5", "sha256", "sha512", "sha3_256", "sha3_512"}
	for _, hashAlgorithim := range hashAlgorithims {
		for _, password := range passwords {
			hashedPassword := getHash(hashAlgorithim, password)
			if !checkFoundPasswords(foundPasswords, hashedPassword) {
				if hashedPassword == hash {
					foundPassword(password, hashedPassword, hashAlgorithim, foundPasswords, foundPasswordsChannel)
				}
			}
		}
	}
}
