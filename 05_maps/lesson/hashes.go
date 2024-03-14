package main

import (
	"fmt"
	"golang.org/x/crypto/sha3"
)

func hashes() {
	inputData := []string{
		"Dior Hawkins",
		"Victor Chan",
		"Hattie Leblanc",
		"Braden Bartlett",
		"Aubrielle Schmidt",
	}
	// naive hash functions
	for _, in := range inputData {
		fmt.Printf("Naive (first char): %s -> %s\n", in, naiveHashByFirstCharHash(in))
	}

	fmt.Println("--------------------------")
	//
	//for _, in := range inputData {
	//	start := time.Now()
	//	fmt.Printf("Naive length (took: %v): %s -> %d\n", time.Since(start), in, naiveHashByLength(in))
	//}
	//
	//fmt.Println("--------------------------")
	//
	//for _, in := range inputData {
	//	start := time.Now()
	//	h := hashSha3(in)
	//	fmt.Printf("SHA3 (took: %v): %s -> %s\n", time.Since(start), in, h)
	//}

	//strong hash (but slow, not suited for hashmaps)
	//userPassword := "mySecretPassword"
	//start := time.Now()
	//
	//hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userPassword), bcrypt.DefaultCost)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//fmt.Printf("bcrypt hash from %s calculated in %v: %s", userPassword, time.Since(start), hashedPassword)
}

func naiveHashByFirstCharHash(in string) string {
	if len(in) > 0 {
		return in[:1]
	}

	return ""
}

func naiveHashByLength(in string) uint8 {
	return uint8(len(in) % 255)
}

func hashSha3(in string) string {
	hash := sha3.New512()
	_, _ = hash.Write([]byte(in))

	// Get the resulting encoded byte slice
	sha3 := hash.Sum(nil)

	// Convert the encoded byte slice to a string
	return fmt.Sprintf("%x", sha3)
}
