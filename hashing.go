package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"

	"golang.org/x/crypto/sha3"
)

func sha1Hash(plaintext string) string {
	hasher := sha1.New()
	return hashText(hasher, plaintext)
}

func md5Hash(plaintext string) string {
	hasher := md5.New()
	return hashText(hasher, plaintext)
}

func sha256Hash(plaintext string) string {
	hasher := sha256.New()
	return hashText(hasher, plaintext)
}

func sha512Hash(plaintext string) string {
	hasher := sha512.New()
	return hashText(hasher, plaintext)
}

func sha3_256Hash(plaintext string) string {
	hasher := sha3.New256()
	return hashText(hasher, plaintext)
}

func sha3_512Hash(plaintext string) string {
	hasher := sha3.New512()
	return hashText(hasher, plaintext)
}
