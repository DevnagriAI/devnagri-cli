package main

import (
	"encoding/base64"
	"fmt"
)

func base64Encode(src []byte) []byte {
	return []byte(base64.StdEncoding.EncodeToString(src))
}

func base64Decode(src []byte) ([]byte, error) {
	return base64.StdEncoding.DecodeString(string(src))
}

func main() {
	// encode

	hello := "PHJlc291cmNlcyB0b29sczppZ25vcmU9IkV4dHJhVHJhbnNsYXRpb24iIHhtbG5zOnRvb2xzPSJodHRwOi8vc2NoZW1hcy5hbmRyb2lkLmNvbS90b29scyI+CiAgICA8c3RyaW5nIG5hbWU9ImFwcF9uYW1lIj48L3N0cmluZz4KICAgIDxzdHJpbmcgbmFtZT0iaGludF9hY3R1YWwiPjwvc3RyaW5nPgogPC9yZXNvdXJjZXM+"

	debyte := base64Encode([]byte(hello))
	fmt.Println(debyte)
	// decode
	enbyte, err := base64Decode(debyte)
	if err != nil {
		fmt.Println(err.Error())
	}

	if hello != string(enbyte) {
		fmt.Println("hello is not equal to enbyte")
	}

	fmt.Println(string(enbyte))
}
