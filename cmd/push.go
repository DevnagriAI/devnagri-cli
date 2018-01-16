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
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	//"github.com/FourtekIT/devnagri-cli/utils"
	"github.com/FourtekIT/devnagri-cli/config"
	"github.com/Jeffail/gabs"
	"github.com/spf13/cobra"
	"gopkg.in/resty.v1"
)

// pushCmd represents the push command
var pushCmd = &cobra.Command{
	Use:   "push",
	Short: "This command pushes the untranslated files from Devnagri",
	Long:  `This command transfers all the untranslated local files to the Devnagri platform on a language basis.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("push called")
	},
}

func init() {
	rootCmd.AddCommand(pushCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// pushCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// pushCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func listAllFilesAndPush() {
	var ClientID = config.FetchAndValidate("ClientID") // returns string

	var ClientSecret = config.FetchAndValidate("ClientSecret") // returns string

	var ProjectKey = config.FetchAndValidate("ProjectKey") // returns string

	var AccessToken = config.FetchAndValidate("AccessToken") // returns string

	filename := "./en/CallingPapaPro2.xml"

	resp, err := resty.R().
		SetHeader("Accept", "application/json").
		SetHeader("Content-Type", "multipart/form-data").
		SetAuthToken("eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiIsImp0aSI6IjE0MmI1YjlkOTI3ZGQ2ZjI3ODcyZDJmNDFmNWQzNjhkOGU3MmVhZjk2YzE2YTdmMTZjM2ZkM2I5ZmVmNjJiMDFmMTBiOWQwMTMxNzJjYTIyIn0.eyJhdWQiOiIxIiwianRpIjoiMTQyYjViOWQ5MjdkZDZmMjc4NzJkMmY0MWY1ZDM2OGQ4ZTcyZWFmOTZjMTZhN2YxNmMzZmQzYjlmZWY2MmIwMWYxMGI5ZDAxMzE3MmNhMjIiLCJpYXQiOjE1MTU1OTA4NDksIm5iZiI6MTUxNTU5MDg0OSwiZXhwIjoxNTQ3MTI2ODQ5LCJzdWIiOiIxIiwic2NvcGVzIjpbXX0.inJAzB_9Vjx9YXHQgZYCy3HxbTaSV2VSjKYgtB_s4ZrrS9vucEnEF4Ijh8wPMylxDLI7e8iGfRObK0rGTBqnCw94ElcmeYPM7GfWRqWW5Bz8ctJhXH7-VGml22fT6ahFVHssbVcx636xwPlydhN4fBay0TElXgC5QkSmzINFPMF3FRJ54WaAsTAvwpsnAIUfigZdNggSjMnROp-i7dtF5M9Ty3utLM2kmlEZVC9lhhLUvJujvbWJwCGDOStrBv4iapWQ4qPZ8AuaVTmigW_hVinkZ6GGTp8GT07GGF9iwDjPnSyCzEx7Ta4S0U3nSeJVogxmsO7N1hZrmCuZmUtEkSPMSP0wIhrMrranpzqFRqvsYve6hp8BBE9qxbuiEwc3F-2xFm-R6esmUu5mpWcRlM1imenuvb9C-lotKf0znM9neVimWhMe7OqENcGXTx7puTyZoK0wUmrn-2_vIxfwtU5s9aBomlw6qGQt6k14pgKA69bp2C51C46HyKa_QrEvol-_3YvccjvR3hO-jcAj_SLALYk6HJp1eSqy_tGZZ5QW4-4QvUcPaa3E4sXA4Ye_4OB2dGRcKVuhR5jzpaYJC-_qGRY0dhfYdpKjEgc1o3pitzpNW0iWzImN9iidRThu_YFP_RwPk2Qragch55f086JNdM7jGPkjcEUcgGIs3PE").
		SetFile("file[0][file]", filename).
		SetFormData(map[string]string{
			"client_id":          ClientID,
			"client_secret":      ClientSecret,
			"project_key":        ProjectKey,
			"file[0][hash]":      sha256Hash(filename),
			"file[0][extension]": "xml",
			"file[0][file_type]": "xml",
			"file[0][location]":  filename,
		}).
		Post("http://dev.devnagri.co.in/api/project/push")

	if err != nil {
		panic(err)
	}

	fmt.Println(resp)
	fmt.Println("\n\n")

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
	//fmt.Println("Reading content.txt")
	fileContent := decodeBase64(string(dat))
	fmt.Println(fileContent)

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
