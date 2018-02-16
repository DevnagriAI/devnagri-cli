package main

import (
	"encoding/base64"
	"fmt"
	//	"io/ioutil"
	"log"
	"os"
	//	"reflect"

	"github.com/FourtekIT/devnagri-cli/config"
	"github.com/Jeffail/gabs"
	"github.com/spf13/viper"
	"gopkg.in/resty.v1"
)

func main() {
	saveResponseAndConvert()
}

func decodeBase64(cypher string) string {

	data, _ := base64.StdEncoding.DecodeString(cypher)

	return string(data)

}

func saveResponseAndConvert() {

	var ClientID = config.FetchAndValidate("ClientID") // returns string

	var ClientSecret = config.FetchAndValidate("ClientSecret") // returns string

	var ProjectKey = config.FetchAndValidate("ProjectKey") // returns string

	var AccessToken = config.FetchAndValidate("AccessToken") // returns string

	resp, err := resty.R().
		SetHeader("Accept", "application/json").
		SetHeader("Content-Type", "multipart/form-data").
		SetAuthToken(AccessToken).
		SetFormData(map[string]string{
			"client_id":     ClientID,
			"client_secret": ClientSecret,
			"project_key":   ProjectKey}).
		Post("http://dev.devnagri.co.in/api/project/pull")

	if err != nil {
		panic(err)
	}

	//fmt.Println(resp)
	resJson, _ := gabs.ParseJSON([]byte(resp.String()))

	fmt.Println(resJson)
	children, _ := resJson.S("file_content").Children()
	child := children[0]
	//fmt.Println(child.String())

	//TODO: Iterate this over all the file names recieved from the remote
	file, err := os.Create("temp.txt")
	if err != nil {
		log.Fatal("Cannot create file", err)
	}
	defer file.Close()

	_, err = file.WriteString(child.String())

	// This works perfectly
	//fmt.Println("Decoding the string manually")
	//x := "PCEtLSBUcmFuc2xhdGVkIEJ5IERldm5hZ3JpIC0tPgo8IS0tIGh0dHA6Ly9kZXZuYWdyaS5jb20gLS0+CjxyZXNvdXJjZXMgdG9vbHM6aWdub3JlPSJFeHRyYVRyYW5zbGF0aW9uIiB4bWxuczp0b29scz0iaHR0cDovL3NjaGVtYXMuYW5kcm9pZC5jb20vdG9vbHMiPgogICAgPHN0cmluZyBuYW1lPSJhcHBfbmFtZSI+PC9zdHJpbmc+CiAgICA8c3RyaW5nIG5hbWU9ImhpbnRfYWN0dWFsIj48L3N0cmluZz4KIDwvcmVzb3VyY2VzPg=="
	//decodeBase64(x)

	//dat, _ := ioutil.ReadFile("temp.txt")
	//datString := string(dat)
	//fmt.Println("String Length : ", len(datString))

	//content := datString[1:(len(datString) - 1)]
	//fmt.Println(content)

	//fileContent := decodeBase64(content)
	//fmt.Println(fileContent)

	//fmt.Println("<<< Reading temp file now >>>")
	//fileContent := decodeBase64(string(dat))
	//fmt.Println(fileContent)

	//TODO: This should be abstracted to various languages

	//TODO: Read all the languages mentioned in the YAML file

	fieldValue := viper.GetStringSlice("TargetLanguages") // returns string
	fmt.Println(fieldValue)

	/*
		//TODO: Make the << hi >> dir
		if _, err := os.Stat("./values-hi"); os.IsNotExist(err) {
			os.Mkdir("./values-hi", os.ModePerm)
		}

		//TODO: Store the content of temp into the actual file

		responseFile, err := os.Create("./values-hi/strings.xml")

		if err != nil {
			log.Fatal("Cannot create file", err)
		}
		defer responseFile.Close()

		_, err = responseFile.WriteString(fileContent)

		os.Remove("temp.txt")

	*/
	fmt.Println("Done!")

}