package utils

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"io"
	"log"
	"os"
)

func sha256Hash(fileName string) string {

	f, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	h := sha256.New()
	if _, err := io.Copy(h, f); err != nil {
		log.Fatal(err)
	}

	//fmt.Printf("%x", h.Sum(nil))
	//return h.Sum(nil)

	hashString := hex.EncodeToString(h.Sum(nil))
	return hashString
}

func decodeBase64(cypher string) string {
	println("\n\nRecieved at decodeBase64 function\n\n")
	println(cypher)
	data, _ := base64.StdEncoding.DecodeString(cypher)

	return string(data)

}
