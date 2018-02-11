package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/FourtekIT/devnagri-cli/config"
	"gopkg.in/resty.v1"
)

func main() {
	var ClientID = config.FetchAndValidate("ClientID") // returns string

	var ClientSecret = config.FetchAndValidate("ClientSecret") // returns string

	var ProjectKey = config.FetchAndValidate("ProjectKey") // returns string

	var AccessToken = config.FetchAndValidate("AccessToken") // returns string

	var RootFolder = config.FetchAndValidate("RootFolder") // returns string

	filename := "./" + RootFolder + "/strings.xml"

	var Extension = config.FetchAndValidate("Extension")

	resp, err := resty.R().
		SetHeader("Accept", "application/json").
		SetHeader("Content-Type", "multipart/form-data").
		SetAuthToken(AccessToken).
		SetFile("file[0][file]", filename).
		SetFormData(map[string]string{
			"client_id":          ClientID,
			"client_secret":      ClientSecret,
			"project_key":        ProjectKey,
			"file[0][hash]":      sha256Hash(filename),
			"file[0][extension]": Extension,
			"file[0][file_type]": "xml",
			"file[0][location]":  filename,
		}).
		Post("http://dev.devnagri.co.in/api/project/push")

	if err != nil {
		panic(err)
	}

	fmt.Println(resp)
}

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
