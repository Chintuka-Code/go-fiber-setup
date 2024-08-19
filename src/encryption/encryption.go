package main

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strings"
)

type PasswordEncryption struct {
	saltLength int
}

func New(saltLength int) (*PasswordEncryption, error) {

	if saltLength <= 0 {
		return nil, fmt.Errorf("invalid salt length: %d. Must be greater than 0", saltLength)
	}

	return &PasswordEncryption{saltLength: saltLength}, nil
}

func (pe *PasswordEncryption) GenerateSalt() ([]byte, error) {
	salt := make([]byte, pe.saltLength)
	_, err := rand.Read(salt)

	if err != nil {
		return nil, err
	}
	return salt, nil
}

func (pe *PasswordEncryption) GenerateHash(plainText string, salt []byte) string {
	hashFnc := sha256.New()
	hashFnc.Write([]byte(plainText))
	hashFnc.Write(salt)
	hashBytes := hashFnc.Sum(nil)
	hashString := hex.EncodeToString(hashBytes)
	saltStr := hex.EncodeToString(salt)
	return fmt.Sprintf("%s$%s", saltStr, hashString)
}

func (pe *PasswordEncryption) VerifyHash(storedHash, plaintextPassword string) bool {
	parts := strings.Split(storedHash, "$")
	if len(parts) != 2 {
		return false
	}

	saltHex := parts[0]
	storedHashHex := parts[1]
	salt, err := hex.DecodeString(saltHex)

	if err != nil {
		return false
	}

	computedHash := pe.GenerateHash(plaintextPassword, salt)
	return computedHash == storedHashHex
}

func main() {

	pe, peErr := New(15)
	if peErr != nil {
		fmt.Print(peErr)
	}

	salt, err := pe.GenerateSalt()
	if err != nil {
		fmt.Println("Something went wrong")
	}
	hash := pe.GenerateHash("Hello", salt)
	fmt.Println(hash)

}
