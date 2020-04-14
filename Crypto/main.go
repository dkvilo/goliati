package crypto

import (
	"crypto/sha256"
	"log"
	"encoding/hex"
)

// GenerateHash func creates hush based on data passed in
func GenerateHash(data ...string) string {
	hash := sha256.New()
	hex.EncodeToString(hash.Sum(nil))

	var tempString string
	for strIndex := range data {
		tempString += string(data[strIndex]) // cast var if isn't already string
	}
	_, err := hash.Write([]byte(tempString))
	if err != nil {
		log.Fatal(err)
		recover()
	}

	return hex.EncodeToString(hash.Sum(nil))
}
