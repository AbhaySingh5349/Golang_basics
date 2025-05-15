package main

import (
	"crypto/rand"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"fmt"
	"io"
)

// for cryptographically secured random nums, we don't need seeding as "rand num generator" is designed to be secured & unpredicatble

func main() {

	password := "password123"

	hash256 := sha256.Sum256([]byte(password))
	hash512 := sha512.Sum512([]byte(password)) // string -> byte slice

	fmt.Println("password: ", password)
	fmt.Println("hash: ", hash256)
	fmt.Println("hash512: ", hash512)
	fmt.Printf("SHA-256 Hash hex val: %x\n", hash256)
	fmt.Printf("SHA-512 Hash hex val: %x\n", hash512)

	// "salt" adds 2nd layer of security by adding unique value
	salt, err := generateSalt()
	fmt.Println("Original Salt:", salt)
	fmt.Printf("Original Salt hex val: %x\n", salt)
	if err != nil {
		fmt.Println("Error generating salt:", err)
		return
	}

	// Hash the password with salt
	signUpHash := hashPassword(password, salt)

	// Store the salt and password in db
	saltStr := base64.StdEncoding.EncodeToString(salt)
	fmt.Println("Base64 converted Salt:", saltStr) // simulate as storing in db
	fmt.Println("Hash:", signUpHash)               // simulate as storing in db
	hashOriginalPassword := sha256.Sum256([]byte(password))
	fmt.Println("Hash of just the password string without salt:", base64.StdEncoding.EncodeToString(hashOriginalPassword[:]))

	// verify
	// retrieve the saltStr and decode it
	decodedSalt, err := base64.StdEncoding.DecodeString(saltStr)
	if err != nil {
		fmt.Println("Unable to decode salt:", err)
		return
	}
	loginHash := hashPassword("password124", decodedSalt)

	// Compare the stored signUpHash with loginHash
	if signUpHash == loginHash {
		fmt.Println("Password is correct. You are logged in.")
	} else {
		fmt.Println("Login failed. Please check user credentials.")
	}
}

func generateSalt() ([]byte, error) {
	salt := make([]byte, 16)                 // 16 size byte slice
	_, err := io.ReadFull(rand.Reader, salt) // reads len of byte slice from "reader" into "salt slice"
	if err != nil {
		return nil, err
	}
	return salt, nil
}

// Function to hash password
func hashPassword(password string, salt []byte) string {
	saltedPassword := append(salt, []byte(password)...) // append "password byte slice to salt"
	hash := sha256.Sum256(saltedPassword)
	return base64.StdEncoding.EncodeToString(hash[:])
}
