package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	str := base64.StdEncoding.EncodeToString([]byte("Hello, playground"))
	fmt.Println(str)

	cypher := "PHJlc291cmNlcyB0b29sczppZ25vcmU9IkV4dHJhVHJhbnNsYXRpb24iIHhtbG5zOnRvb2xzPSJodHRwOi8vc2NoZW1hcy5hbmRyb2lkLmNvbS90b29scyI+CiAgICA8c3RyaW5nIG5hbWU9ImFwcF9uYW1lIj48L3N0cmluZz4KICAgIDxzdHJpbmcgbmFtZT0iaGludF9hY3R1YWwiPjwvc3RyaW5nPgogPC9yZXNvdXJjZXM+"

	decodeBase64(cypher)

}

func decodeBase64(cypher string) string {

	data, _ := base64.StdEncoding.DecodeString(cypher)

	fmt.Printf("%q\n", data)
	return string(data)

}
