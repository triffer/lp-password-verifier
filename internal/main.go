package internal

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"strings"
)

func Main() {
	var verifiedPwnedPasswords []PasswordRecord

	// We want to skip the first line because it contains the header
	for _, record := range LoadPasswordRecords()[1:] {
		fmt.Printf("Processing %s\n", record.Name)

		passwordHash := getSha1Hash(record.Password)
		pwnedPasswords := FindPwnedPasswordsByHash(passwordHash)

		if isPasswordPwned(passwordHash, pwnedPasswords) {
			verifiedPwnedPasswords = append(verifiedPwnedPasswords, record)
		}
	}

	fmt.Printf("\nThe following passwords have been pwned\n\n")
	for _, pwnedPassword := range verifiedPwnedPasswords {
		fmt.Printf("Name: %s \nPassword: %s \n", pwnedPassword.Name, pwnedPassword.Password)
		fmt.Println("--------------------------------")
	}
}

func getSha1Hash(text string) string {
	hasher := sha1.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}

func isPasswordPwned(passwordHash string, pwnedPasswordHashes []string) bool {
	for _, pwnedHash := range pwnedPasswordHashes {
		if strings.EqualFold(passwordHash, pwnedHash) {
			return true
		}
	}

	return false
}

