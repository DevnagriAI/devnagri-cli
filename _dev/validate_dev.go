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

package main

import (
	"fmt"
	//"github.com/FourtekIT/devnagri-cli/config"
	"os"

	"github.com/Jeffail/gabs"
	"gopkg.in/resty.v1"
)

func main() {

	resp, err := resty.R().
		SetHeader("Content-Type", "multipart/form-data").
		SetFormData(map[string]string{
			"client_id":     "3",
			"client_secret": "3WnUqVSP7Vhs8DU7FInIrwHIVMg9twGshcpswlJW",
			"project_key":   "2e12635aca73c7d39ec76a514d7490a6"}).
		Post("http://dev.devnagri.co.in/api/key/validations")

	if err != nil {
		panic(err)
	}

	//TODO: Save the returned access_token to the .devnagri.yaml
	// We do this by appending the contents of the .devnagri.yaml file

	jsonParsed, _ := gabs.ParseJSON([]byte(resp.String()))
	accessToken := jsonParsed.Path("access_token").Data()
	//fmt.Println(access_token)
	filename := "./.devnagri.yaml"

	f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}

	defer f.Close()
	accessTokenString := "AccessToken: " + accessToken.(string)
	fmt.Println(accessTokenString)

	f.WriteString(accessTokenString)

	if err != nil {
		panic(err)
	}
}
