package main

var hashAlgorithms = map[string]func(string) string{
	"sha1":     sha1Hash,
	"md5":      md5Hash,
	"sha256":   sha256Hash,
	"sha512":   sha512Hash,
	"sha3_256": sha3_256Hash,
	"sha3_512": sha3_512Hash,
}

func getHash(hashType string, plaintext string) string {
	// note: a map function is faster than a switch
	// source: https://hashrocket.com/blog/posts/switch-vs-map-which-is-the-better-way-to-branch-in-go
	for _, hashAlgorithm := range hashAlgorithms {
		return hashAlgorithm(plaintext)
	}
	panic("Hash type not supported!")
}

func checkFoundPasswords(foundPasswords *[]string, hashedPassword string) bool {
	for _, foundPassword := range *foundPasswords {
		if foundPassword == hashedPassword {
			return true
		}
	}
	return false
}

func foundPassword(password string, hashedPassword string, hashAlgorithim string, foundPasswords *[]string) {
	// fmt.Printf("Matched password (%s): %s, %s\n", hashAlgorithim, password, hashedPassword)
	writeCSV(resultsFile, []string{password, hashedPassword, hashAlgorithim})
	*foundPasswords = append(*foundPasswords, hashedPassword)
}

func checkPassword(passwords []string, foundPasswords *[]string, hash string) {
	// for each hash algorithm, check all passwords
	hashAlgorithims := []string{"sha1", "md5", "sha256", "sha512", "sha3_256", "sha3_512"}
	for _, hashAlgorithim := range hashAlgorithims {
		for _, password := range passwords {
			hashedPassword := getHash(hashAlgorithim, password)
			if !checkFoundPasswords(foundPasswords, hashedPassword) {
				if hashedPassword == hash {
					foundPassword(password, hashedPassword, hashAlgorithim, foundPasswords)
				}
			}
		}
	}
}