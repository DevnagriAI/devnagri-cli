// Copyright © 2017 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/FourtekIT/devnagri-cli/config"
	"github.com/Jeffail/gabs"
	"github.com/spf13/cobra"
	//	"reflect"
	"gopkg.in/resty.v1"
)

// pullCmd represents the pull command
var pullCmd = &cobra.Command{
	Use:   "pull",
	Short: "This command pulls the translated files from Devnagri",
	Long:  `When all the files for a language have been translated, they can be pulled from the Devnagri platform to the local filesystem using the CLI tool.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("pull called")
		saveResponseAndConvert()
	},
}

//TODO: save the response file_content of the request into a file and then read-decode-save it again.
func init() {
	rootCmd.AddCommand(pullCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// pullCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// pullCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func saveResponseAndConvert() {
	var ClientID = config.FetchAndValidate("ClientID") // returns string

	var ClientSecret = config.FetchAndValidate("ClientSecret") // returns string

	var ProjectKey = config.FetchAndValidate("ProjectKey") // returns string

	var AccessToken = config.FetchAndValidate("AccessToken") // returns string

	resp, err := resty.R().
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

	//	fmt.Println(resp)

	resJson, _ := gabs.ParseJSON([]byte(resp.String()))
	children, _ := resJson.S("file_content").Children()
	child := children[0]

	fmt.Println(child.String())

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

	dat, _ := ioutil.ReadFile("temp.txt")
	datString := string(dat)
	//fmt.Println("String Length : ", len(datString))
	content := datString[1:(len(datString) - 1)]
	//fmt.Println(content)
	fileContent := decodeBase64(content)
	//fmt.Println(fileContent)
	//fmt.Println("<<< Reading temp file now >>>")
	//fileContent := decodeBase64(string(dat))
	//fmt.Println(fileContent)

	//TODO: Store the content of temp into the actual file
	responseFile, err := os.Create("responseFile.txt")
	if err != nil {
		log.Fatal("Cannot create file", err)
	}
	defer responseFile.Close()

	_, err = responseFile.WriteString(fileContent)

	//TODO: Delete the temp file
}

func decodeBase64(cypher string) string {
	//println("\n\nRecieved at decodeBase64 function\n\n")
	//fmt.Println(cypher)

	data, _ := base64.StdEncoding.DecodeString(cypher)

	//fmt.Println(string(data))

	println("Text after decoding")

	//fmt.Println(data)
	//fmt.Println(string(data))

	return string(data)

}
