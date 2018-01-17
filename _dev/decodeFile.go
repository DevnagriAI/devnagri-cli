package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
)

func main() {

	dat, _ := ioutil.ReadFile("temp.txt")
	datString := string(dat)
	//fmt.Println("String Length : ", len(datString))
	content := datString[1:(len(datString) - 1)]
	//fmt.Println(content)
	fileContent := decodeBase64(content)
	//fmt.Println(fileContent)
	//fmt.Println("<<< Reading temp file now >>>")
	//fileContent := decodeBase64(string(dat))
	fmt.Println(fileContent)
}

func decodeBase64(cypher string) string {

	data, _ := base64.StdEncoding.DecodeString(cypher)

	stringData := string(data)

	return stringData
}
